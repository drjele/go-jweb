package module

import (
    `gitlab.com/drjele-go/jweb/config/parameter`
    `gitlab.com/drjele-go/jweb/kernel`
)

type List []Module
type Map map[string]Module

type Module interface {
    GetName() string

    ConfigurationRequired() bool

    Boot(kernel *kernel.Kernel, config *parameter.Yaml)
}
