package jwebresponse

const (
    ContentTypeCsv = `text/csv`
)

type Response interface {
    /** @todo add set payload to this */

    Render()

    GetHttpStatus() int

    SetHttpStatus(httpStatus int)
}
