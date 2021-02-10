package response

import (
    `github.com/gin-gonic/gin`
)

type DataPayload []byte

func NewData(context *gin.Context, contentType string) *Data {
    return &Data{
        context:     context,
        contentType: contentType,
    }
}

type Data struct {
    context     *gin.Context
    contentType string
    httpStatus  int
    payload     DataPayload
}

func (d *Data) Render() {
    (*d.context).Data(d.httpStatus, d.contentType, d.payload)
}

func (d *Data) GetHttpStatus() int {
    return d.httpStatus
}

func (d *Data) SetHttpStatus(httpStatus int) {
    d.httpStatus = httpStatus
}

func (d *Data) GetPayload() DataPayload {
    return d.payload
}

func (d *Data) SetPayload(payload DataPayload) {
    d.payload = payload
}
