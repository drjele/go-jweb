package jwebkernel

import (
    jwebconfig `gitlab.com/drjele-go/jweb/kernel/config`
    jwebenvironment `gitlab.com/drjele-go/jweb/kernel/environment`
    jwebflag `gitlab.com/drjele-go/jweb/kernel/flag`
)

func New() *Kernel {
    kernel := Kernel{}

    kernel.environment = jwebenvironment.New()

    kernel.config = jwebconfig.New(kernel.environment)

    kernel.flags = jwebflag.New(kernel.environment)

    return &kernel
}

type Kernel struct {
    environment *jwebenvironment.Environment
    config      *jwebconfig.Config
    flags       *jwebflag.Flag
}

func (k *Kernel) GetEnvironment() *jwebenvironment.Environment {
    return k.environment
}

func (k *Kernel) GetConfig() *jwebconfig.Config {
    return k.config
}

func (k *Kernel) GetFlags() *jwebflag.Flag {
    return k.flags
}
