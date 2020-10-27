package jweb

import (
    `flag`
    `fmt`

    jwebcli `gitlab.com/drjele-go/jweb/cli`
    jwebcommand `gitlab.com/drjele-go/jweb/cli/command`
    jwebflag `gitlab.com/drjele-go/jweb/cli/flag`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebhttp `gitlab.com/drjele-go/jweb/http`
    jwebroute `gitlab.com/drjele-go/jweb/http/routing/route`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
)

func New(
    routeList jwebroute.List,
    commandList jwebcommand.List,
) *Jweb {
    kernel := jwebkernel.New()

    jweb := Jweb{
        kernel:      kernel,
        routeList:   routeList,
        commandList: commandList,
    }

    return &jweb
}

type Jweb struct {
    kernel      *jwebkernel.Kernel
    routeList   jwebroute.List
    commandList jwebcommand.List
}

func (j *Jweb) Run() {
    defer j.handleError()

    switch j.kernel.GetFlags().GetMode() {
    case jwebflag.ModeWeb:
        jwebhttp.Run(j.kernel.GetConfig(), j.routeList)
        break
    case jwebflag.ModeCli:
        jwebcli.Run(j.kernel.GetConfig().GetCli(), j.commandList)
        break
    default:
        flag.PrintDefaults()

        jweberror.Fatal(jweberror.New(`invalid application mode "%v"`, j.kernel.GetFlags().GetMode()))
    }
}

func (Jweb) handleError() {
    r := recover()

    if r == nil {
        return
    }

    var err error

    if e, ok := r.(jweberror.Error); ok {
        err = &e
    } else {
        err = fmt.Errorf(`%v`, r)
    }

    jweberror.Fatal(err)
}
