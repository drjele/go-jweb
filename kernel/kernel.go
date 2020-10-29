package jwebkernel

import (
    jwebconfig `gitlab.com/drjele-go/jweb/kernel/config`
    jwebenvironment `gitlab.com/drjele-go/jweb/kernel/environment`
    jwebflag `gitlab.com/drjele-go/jweb/kernel/flag`
)

func New(rootDir string) *Kernel {
    kernel := Kernel{}

    kernel.rootDir = rootDir

    kernel.environment = jwebenvironment.New(rootDir)

    kernel.config = jwebconfig.New(kernel.environment)

    kernel.flags = jwebflag.New(kernel.environment)

    return &kernel
}

type Kernel struct {
    rootDir     string
    environment *jwebenvironment.Environment
    config      *jwebconfig.Config
    flags       *jwebflag.Flag
}

func (k *Kernel) GetRootDir() string {
    return k.rootDir
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
