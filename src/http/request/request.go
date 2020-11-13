package jwebrequest

import (
    `github.com/gin-gonic/gin`

    jweberror `gitlab.com/drjele-go/jweb/src/error`
)

func New(context *gin.Context) *Request {
    return &Request{
        context: context,
    }
}

type Request struct {
    context *gin.Context
}

func (r *Request) GetQueryParams(params map[string]string) map[string]string {
    /** @todo use a struct */

    for param, defaultValue := range params {
        params[param] = (*r.context).DefaultQuery(param, defaultValue)
    }

    return params
}

func (r *Request) GetPostParams(params map[string]string) map[string]string {
    /** @todo use a struct */

    for param, defaultValue := range params {
        params[param] = (*r.context).DefaultPostForm(param, defaultValue)
    }

    return params
}

func (r *Request) GetJsonParams(json interface{}) {
    err := r.context.BindJSON(json)

    jweberror.Panic(err)
}
