package error

import (
    `log`
    `runtime/debug`
)

/** @todo maybe also add an error handling module */

/** @deprecated will add a logger module */
func Panic(err error) {
    if err == nil {
        return
    }

    panic(err)
}

/** @deprecated will add a logger module */
func Fatal(err error) {
    if err == nil {
        return
    }

    /** @todo only in dev mode */
    debug.PrintStack()

    log.Fatal(err)
}
