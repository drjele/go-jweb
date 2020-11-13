package flag

import (
    `flag`
    `strings`

    `gitlab.com/drjele-go/jweb/src/kernel/environment`
)

const (
    FlagMode = `mode`
)

func New(environment *environment.Environment) *Flag {
    f := Flag{}

    f.parseFlags(environment)

    return &f
}

type Flag struct {
    mode string
}

func (f *Flag) GetMode() string {
    return f.mode
}

func (f *Flag) parseFlags(e *environment.Environment) {
    flag.StringVar(
        &f.mode,
        FlagMode,
        e.GetDefaultMode(),
        `The mode in which the application should run: `+strings.Join([]string{environment.ModeHttp, environment.ModeCli}, ` | `),
    )

    flag.Parse()
}
