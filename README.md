# Jweb

This is a **Symfony**(https://symfony.com/) inspired framework, written in Go.

It is a work in progress and any contributions are welcomed.

## Todo
* module mechanism
* yaml config files and separate config files for each module
* module config structure
* maybe move database module to different repository
* resolve **config.yaml** params based on **.env** params

## Usage
* the **doc** folder contains **.dist** files with example configurations

```go
package main

import (
    `time`

    `github.com/gin-gonic/gin`
    `gitlab.com/drjele-go/jweb`
    command `gitlab.com/drjele-go/jweb/cli/command`
    jwebresponse `gitlab.com/drjele-go/jweb/http/response`
    route `gitlab.com/drjele-go/jweb/http/routing/route`
)

func main() {
    var commandList command.List

    j := jweb.New(getRouteList(), commandList)

    j.Run()
}

func getRouteList() route.List {
    routeList := route.List{}

    routeList = append(
        routeList,
        route.New(
            route.MethodGet,
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
