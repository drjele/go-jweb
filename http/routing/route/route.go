package jwebroute

import (
    `github.com/gin-gonic/gin`

    jwebresponse `gitlab.com/drjele-go/jweb/http/response`
)

const (
    MethodGet    = `GET`
    MethodPost   = `POST`
    MethodPut    = `PUT`
    MethodPatch  = `PATCH`
    MethodDelete = `DELETE`
)

type List []*Route
type Map map[string]*Route
type Handler func(*gin.Context) jwebresponse.Response

func New(method string, path string, handler Handler) *Route {
    return &Route{
        method:  method,
        path:    path,
        handler: handler,
    }
}

type Route struct {
    method  string
    path    string
    handler Handler
}

func (r *Route) GetMethod() string {
    return r.method
}

func (r *Route) GetPath() string {
    return r.path
}

func (r *Route) GetHandler() Handler {
    return r.handler
}
