package jwebcommand

import (
    `github.com/urfave/cli/v2`
)

type List []Command

type Command interface {
    GetName() string

    GetDescription() string

    GetFlags() []cli.Flag

    Execute(context *cli.Context) error
}
