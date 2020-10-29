package jwebmodule

import (
    jwebparameter `gitlab.com/drjele-go/jweb/config/parameter`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
)

type List []Module
type Map map[string]Module

type Module interface {
    GetName() string

    ConfigurationRequired() bool

    Validate(config *jwebparameter.Yaml) error

    Boot(kernel *jwebkernel.Kernel, config *jwebparameter.Yaml)
}
