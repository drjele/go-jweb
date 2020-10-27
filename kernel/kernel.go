package jwebkernel

import (
    jwebflag `gitlab.com/drjele-go/jweb/cli/flag`
    jwebconfig `gitlab.com/drjele-go/jweb/kernel/config`
)

func New() *Kernel {
    kernel := Kernel{}

    kernel.config = jwebconfig.New()

    kernel.flags = jwebflag.New()

    return &kernel
}

type Kernel struct {
    config *jwebconfig.Config
    flags  *jwebflag.Flag
}

func (k *Kernel) GetConfig() *jwebconfig.Config {
    return k.config
}

func (k *Kernel) GetFlags() *jwebflag.Flag {
    return k.flags
}

func (k *Kernel) GetEnv() string {
    return k.config.GetEnv()
}
