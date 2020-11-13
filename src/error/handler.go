package jweberror

import (
    `log`
    `runtime/debug`
)

func NewHandler(env string) *Handler {
    handler := Handler{}

    handler.env = env

    return &handler
}

type Handler struct {
    env string
}

func (h *Handler) Panic(err error) {
    if err == nil {
        return
    }

    panic(err)
}

func (h *Handler) Fatal(err error) {
    if err == nil {
        return
    }

    /** @todo only in dev mode */
    debug.PrintStack()

    log.Fatal(err)
}
