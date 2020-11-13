package jwebmodule

import (
    parameter `gitlab.com/drjele-go/jweb/src/config/parameter`
    jwebkernel `gitlab.com/drjele-go/jweb/src/kernel`
)

type List []Module
type Map map[string]Module

type Module interface {
    GetName() string

    ConfigurationRequired() bool

    Boot(kernel *jwebkernel.Kernel, config *parameter.Yaml)
}
