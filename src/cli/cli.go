package jwebcli

import (
    `os`

    `github.com/urfave/cli/v2`

    `gitlab.com/drjele-go/jweb/src/cli/command`
    `gitlab.com/drjele-go/jweb/src/cli/config`
    jweberror `gitlab.com/drjele-go/jweb/src/error`
    `gitlab.com/drjele-go/jweb/src/kernel/flag`
    `gitlab.com/drjele-go/jweb/src/utility/slice`
)

type list []*cli.Command

func Run(config *config.Config, commandList command.List) {
    if len(commandList) == 0 {
        jweberror.Fatal(
            jweberror.New(`no commands were set for cli mode`),
        )
    }

    app := &cli.App{
        Name:  config.GetName(),
        Usage: config.GetDescription(),
        Flags: []cli.Flag{
            &cli.StringFlag{Name: flag.FlagMode, Usage: `this is the primary app flag and should not be used for commands`},
        },
        Commands: buildList(commandList),
    }

    err := app.Run(os.Args)
    jweberror.Fatal(err)
}

func buildList(commands command.List) list {
    var list list
    var commandNames []string

    for _, c := range commands {
        if slice.StringInSlice(c.GetName(), commandNames) {
            jweberror.Fatal(
                jweberror.New(`duplicate command name "%v"`, c.GetName()),
            )
        }

        commandNames = append(commandNames, c.GetName())
        list = append(list, buildCommand(c))
    }

    return list
}

func buildCommand(c command.Command) *cli.Command {
    return &cli.Command{
        Name:        c.GetName(),
        Description: c.GetDescription(),
        Flags:       c.GetFlags(),
        Action: func(context *cli.Context) error {
            return c.Execute(context)
        },
    }
}
