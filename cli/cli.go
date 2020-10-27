package jwebcli

import (
    `os`

    `github.com/urfave/cli/v2`

    jwebcommand `gitlab.com/drjele-go/jweb/cli/command`
    jwebconfig `gitlab.com/drjele-go/jweb/cli/config`
    jwebflag `gitlab.com/drjele-go/jweb/cli/flag`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebslice `gitlab.com/drjele-go/jweb/utility/slice`
)

type list []*cli.Command

func Run(config *jwebconfig.Config, commandList jwebcommand.List) {
    app := &cli.App{
        Name:  config.GetName(),
        Usage: config.GetDescription(),
        Flags: []cli.Flag{
            &cli.StringFlag{Name: jwebflag.FlagMode, Usage: `this is the primary app flag and should not be used for commands`},
        },
        Commands: buildList(commandList),
    }

    err := app.Run(os.Args)
    jweberror.Fatal(err)
}

func buildList(commands jwebcommand.List) list {
    var list list
    var commandNames []string

    for _, command := range commands {
        if jwebslice.StringInSlice(command.GetName(), commandNames) {
            jweberror.Fatal(
                jweberror.New(`duplicate command name "%v"`, command.GetName()),
            )
        }

        commandNames = append(commandNames, command.GetName())
        list = append(list, buildCommand(command))
    }

    return list
}

func buildCommand(command jwebcommand.Command) *cli.Command {
    return &cli.Command{
        Name:        command.GetName(),
        Description: command.GetDescription(),
        Flags:       command.GetFlags(),
        Action: func(context *cli.Context) error {
            return command.Execute(context)
        },
    }
}
