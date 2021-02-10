package http

import (
    `net/http`
    `time`

    `golang.org/x/sync/errgroup`

    jweberror `gitlab.com/drjele-go/jweb/error`
    `gitlab.com/drjele-go/jweb/http/routing/route`
    jwebrouter `gitlab.com/drjele-go/jweb/http/routing/router`
    `gitlab.com/drjele-go/jweb/kernel/config`
    `gitlab.com/drjele-go/jweb/kernel/environment`
)

var (
    group errgroup.Group
)

func Run(
    environment *environment.Environment,
    config *config.Config,
    routeList route.List,
) {
    if len(routeList) == 0 {
        jweberror.Fatal(
            jweberror.New(`no routes were set for http mode`),
        )
    }

    router := jwebrouter.New(environment.GetEnv(), routeList)

    server := &http.Server{
        Addr:         config.GetHttp().GetHost(),
        Handler:      router.GetHttpHandler(),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    group.Go(func() error {
        err := server.ListenAndServe()

        if err != nil && err != http.ErrServerClosed {
            jweberror.Fatal(err)
        }

        return err
    })

    if err := group.Wait(); err != nil {
        jweberror.Fatal(err)
    }
}
