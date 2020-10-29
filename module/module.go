package jwebmodule

import (
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
)

type List []Module
type Map map[string]Module

type Module interface {
    GetName() string

    ConfigurationRequired() bool

    Boot(kernel *jwebkernel.Kernel)
}
