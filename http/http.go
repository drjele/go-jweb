package jwebhttp

import (
    `net/http`
    `time`

    `golang.org/x/sync/errgroup`

    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebroute `gitlab.com/drjele-go/jweb/http/routing/route`
    jwebrouter `gitlab.com/drjele-go/jweb/http/routing/router`
    jwebconfig `gitlab.com/drjele-go/jweb/kernel/config`
)

var (
    group errgroup.Group
)

func Run(config *jwebconfig.Config, routeList jwebroute.List) {
    router := jwebrouter.New(config.GetEnv(), routeList)

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
