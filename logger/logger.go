package logger

import (
    jwebparameter `gitlab.com/drjele-go/jweb/config/parameter`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
)

const (
    Name = `logger`
)

func New() *Logger {
    logger := Logger{}

    return &logger
}

type Logger struct {
    kernel *jwebkernel.Kernel
}

func (l *Logger) GetName() string {
    return Name
}

func (l *Logger) ConfigurationRequired() bool {
    return true
}

func (l *Logger) Boot(kernel *jwebkernel.Kernel, yamlConfig *jwebparameter.Yaml) {
    l.kernel = kernel
}
