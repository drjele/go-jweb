package jwebresponse

import (
    `github.com/gin-gonic/gin`
)

type JsonPayload gin.H

func NewJson(context *gin.Context) *Json {
    return &Json{
        context: context,
    }
}

type Json struct {
    context    *gin.Context
    httpStatus int
    payload    JsonPayload
}

func (j *Json) Render() {
    (*j.context).JSON(j.httpStatus, j.payload)
}

func (j *Json) GetHttpStatus() int {
    return j.httpStatus
}

func (j *Json) SetHttpStatus(httpStatus int) {
    j.httpStatus = httpStatus
}

func (j *Json) GetPayload() JsonPayload {
    return j.payload
}

func (j *Json) SetPayload(payload JsonPayload) {
    j.payload = payload
}
