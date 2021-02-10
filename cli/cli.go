package cli

import (
    `os`

    `github.com/urfave/cli/v2`

    jwebcommand `gitlab.com/drjele-go/jweb/cli/command`
    `gitlab.com/drjele-go/jweb/cli/config`
    jweberror `gitlab.com/drjele-go/jweb/error`
    `gitlab.com/drjele-go/jweb/kernel/flag`
    `gitlab.com/drjele-go/jweb/utility`
)

type list []*cli.Command

func Run(config *config.Config, commandList jwebcommand.List) {
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

func buildList(commands jwebcommand.List) list {
    var list list
    var commandNames []string

    for _, command := range commands {
        if utility.StringInSlice(command.GetName(), commandNames) {
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
