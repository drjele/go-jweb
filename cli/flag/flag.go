package jwebflag

import (
    `flag`
    `strings`
)

const (
    ModeWeb = `web`
    ModeCli = `cli`

    FlagMode = `mode`
)

func New() *Flag {
    f := Flag{}

    f.parseFlags()

    return &f
}

type Flag struct {
    mode string
}

func (f *Flag) GetMode() string {
    return f.mode
}

func (f *Flag) parseFlags() {
    flag.StringVar(
        &f.mode,
        FlagMode,
        ModeWeb,
        `The mode in which the application should run: `+strings.Join([]string{ModeWeb, ModeCli}, ` | `),
    )

    flag.Parse()
}
