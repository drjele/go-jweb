package router

import (
    `fmt`
    `net/http`
    `runtime/debug`
    `time`

    `github.com/gin-gonic/gin`

    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebresponse `gitlab.com/drjele-go/jweb/http/response`
    jwebroute `gitlab.com/drjele-go/jweb/http/routing/route`
    `gitlab.com/drjele-go/jweb/kernel/environment`
)

func New(env string, routeList jwebroute.List) *Router {
    router := Router{
        env:    env,
        routes: jwebroute.Map{},
    }

    for _, route := range routeList {
        path := route.GetPath()
        _, ok := router.routes[path]

        if ok == true {
            jweberror.Fatal(jweberror.New(`multiple routes defined for path "%v"`, path))
        }

        router.routes[path] = route
    }

    return &router
}

type Router struct {
    env    string
    routes jwebroute.Map
}

func (r *Router) GetHttpHandler() http.Handler {
    gin.SetMode(r.getEnv())

    engine := gin.New()
    engine.Use(gin.Recovery())

    for _, route := range r.routes {
        r.attachRoute(engine, route)
    }

    return engine
}

func (r *Router) getEnv() string {
    env := gin.DebugMode

    if r.env == environment.EnvProd {
        env = gin.ReleaseMode
    }

    return env
}

func (r *Router) attachRoute(engine *gin.Engine, route *jwebroute.Route) {
    handler := func(context *gin.Context) {
        r.handleRoute(route, context)
    }

    switch route.GetMethod() {
    case jwebroute.MethodGet:
        engine.GET(route.GetPath(), handler)
        break
    case jwebroute.MethodPost:
        engine.POST(route.GetPath(), handler)
        break
    case jwebroute.MethodPut:
        engine.PUT(route.GetPath(), handler)
        break
    case jwebroute.MethodPatch:
        engine.PATCH(route.GetPath(), handler)
        break
    case jwebroute.MethodDelete:
        engine.DELETE(route.GetPath(), handler)
        break
    default:
        jweberror.Fatal(jweberror.New(`no handler defined for http method "%v"`, route.GetMethod()))
        break
    }
}

func (r *Router) renderResponse(response jwebresponse.Response) {
    response.Render()
}

func (r *Router) handleRoute(route *jwebroute.Route, context *gin.Context) {
    defer r.handleError(context)

    handler := route.GetHandler()
    response := handler(context)

    httpStatus := response.GetHttpStatus()
    if httpStatus == 0 {
        response.SetHttpStatus(http.StatusOK)
    }

    defer r.renderResponse(response)
}

func (r *Router) handleError(context *gin.Context) {
    recoverData := recover()

    if recoverData == nil {
        /** @todo maybe log */
        return
    }

    jsonPayload := jwebresponse.JsonPayload{
        `error`: fmt.Sprintf(`%v`, recoverData),
        `time`:  time.Now(),
    }
    if r.env == environment.EnvDev {
        jsonPayload[`trace`] = string(debug.Stack())
    }

    response := jwebresponse.NewJson(context)

    response.SetHttpStatus(http.StatusInternalServerError)
    response.SetPayload(jsonPayload)

    r.renderResponse(response)
}
