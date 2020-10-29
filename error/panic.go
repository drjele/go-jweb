package jweberror

import (
    `log`
    `runtime/debug`
)

/** @todo add logging */

func Panic(err error) {
    if err == nil {
        return
    }

    panic(err)
}

func Fatal(err error) {
    if err == nil {
        return
    }

    debug.PrintStack()

    log.Fatal(err)
}
