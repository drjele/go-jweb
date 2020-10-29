package jwebflag

import (
    `flag`
    `strings`

    jwebenvironment `gitlab.com/drjele-go/jweb/kernel/environment`
)

const (
    FlagMode = `mode`
)

func New(environment *jwebenvironment.Environment) *Flag {
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

func (f *Flag) parseFlags(environment *jwebenvironment.Environment) {
    flag.StringVar(
        &f.mode,
        FlagMode,
        environment.GetDefaultMode(),
        `The mode in which the application should run: `+strings.Join([]string{jwebenvironment.ModeHttp, jwebenvironment.ModeCli}, ` | `),
    )

    flag.Parse()
}
