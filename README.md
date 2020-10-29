# Jweb

This is a **Symfony**(https://symfony.com/) inspired framework, written in Go.

It is a work in progress and any contributions are welcomed.

## Todo
* module mechanism
* yaml config files and separate config files for each module
* module config structure
* maybe move database module to different repository

## Usage
```go
package main

import (
    `time`

    `github.com/gin-gonic/gin`
    `gitlab.com/drjele-go/jweb`
    jwebcommand `gitlab.com/drjele-go/jweb/cli/command`
    jwebresponse `gitlab.com/drjele-go/jweb/http/response`
    jwebroute `gitlab.com/drjele-go/jweb/http/routing/route`
)

func main() {
    var commandList jwebcommand.List

    j := jweb.New(getRouteList(), commandList)

    j.Run()
}

func getRouteList() jwebroute.List {
    routeList := jwebroute.List{}

    routeList = append(
        routeList,
        jwebroute.New(
            jwebroute.MethodGet,
            `/`,
            func(context *gin.Context) jwebresponse.Response {
                payload := jwebresponse.JsonPayload{`time`: time.Now()}

                response := jwebresponse.NewJson(context)

                response.SetPayload(payload)

                return response
            },
        ),
    )

    return routeList
}
```
