package jweberror

import (
    `log`
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

    log.Fatal(err)
}
