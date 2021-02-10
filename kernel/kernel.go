package kernel

import (
    `gitlab.com/drjele-go/jweb/kernel/config`
    `gitlab.com/drjele-go/jweb/kernel/environment`
    `gitlab.com/drjele-go/jweb/kernel/flag`
)

func New(rootDir string) *Kernel {
    kernel := Kernel{}

    kernel.rootDir = rootDir

    kernel.environment = environment.New(rootDir)

    kernel.config = config.New(kernel.environment)

    kernel.flags = flag.New(kernel.environment)

    return &kernel
}

type Kernel struct {
    rootDir     string
    environment *environment.Environment
    config      *config.Config
    flags       *flag.Flag
}

func (k *Kernel) GetRootDir() string {
    return k.rootDir
}

func (k *Kernel) GetEnvironment() *environment.Environment {
    return k.environment
}

func (k *Kernel) GetConfig() *config.Config {
    return k.config
}

func (k *Kernel) GetFlags() *flag.Flag {
    return k.flags
}
