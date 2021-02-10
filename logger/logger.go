package logger

import (
    `gitlab.com/drjele-go/jweb/config/parameter`
    `gitlab.com/drjele-go/jweb/kernel`
)

const (
    Name = `logger`
)

func New() *Logger {
    logger := Logger{}

    return &logger
}

type Logger struct {
    kernel *kernel.Kernel
}

func (l *Logger) GetName() string {
    return Name
}

func (l *Logger) ConfigurationRequired() bool {
    return true
}

func (l *Logger) Boot(kernel *kernel.Kernel, yamlConfig *parameter.Yaml) {
    l.kernel = kernel
}
