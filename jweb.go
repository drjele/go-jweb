package jweb

import (
    `flag`
    `fmt`
    `log`
    `os`

    jwebcli `gitlab.com/drjele-go/jweb/cli`
    jwebcommand `gitlab.com/drjele-go/jweb/cli/command`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebhttp `gitlab.com/drjele-go/jweb/http`
    jwebroute `gitlab.com/drjele-go/jweb/http/routing/route`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
    jwebenvironment `gitlab.com/drjele-go/jweb/kernel/environment`
    jwebmodule `gitlab.com/drjele-go/jweb/module`
)

func New(
    moduleList jwebmodule.List,
) *Jweb {
    /** @todo maybe split initialization to a boot function */

    dir, err := os.Getwd()
    jweberror.Fatal(err)

    /** @todo maybe not add the trailing slash */
    kernel := jwebkernel.New(dir + `/`)

    jweb := Jweb{
        kernel:      kernel,
        routeList:   jwebroute.List{},
        commandList: jwebcommand.List{},
    }

    jweb.bootModules(moduleList)

    return &jweb
}

type Jweb struct {
    moduleList  jwebmodule.Map
    kernel      *jwebkernel.Kernel
    routeList   jwebroute.List
    commandList jwebcommand.List
}

func (j *Jweb) SetRouteList(routeList jwebroute.List) {
    j.routeList = routeList
}

func (j *Jweb) SetCommandList(commandList jwebcommand.List) {
    j.commandList = commandList
}

func (j *Jweb) Run() {
    defer j.handleError()

    switch j.kernel.GetFlags().GetMode() {
    case jwebenvironment.ModeHttp:
        jwebhttp.Run(
            j.kernel.GetEnvironment(),
            j.kernel.GetConfig(),
            j.routeList,
        )
        break
    case jwebenvironment.ModeCli:
        jwebcli.Run(j.kernel.GetConfig().GetCli(), j.commandList)
        break
    default:
        flag.PrintDefaults()

        jweberror.Fatal(jweberror.New(`invalid application mode "%v"`, j.kernel.GetFlags().GetMode()))
    }
}

func (j *Jweb) GetKernel() *jwebkernel.Kernel {
    return j.kernel
}

func (j *Jweb) GetModule(name string) jwebmodule.Module {
    module, ok := j.moduleList[name]
    if ok == false {
        jweberror.Fatal(jweberror.New(`no module registered for name "%v"`, name))
    }

    return module
}

func (j *Jweb) bootModules(moduleList jwebmodule.List) {
    j.moduleList = jwebmodule.Map{}

    for _, module := range moduleList {
        log.Printf(`boot start "%v"`+"\n", module.GetName())

        j.bootModule(module)

        log.Printf(`boot end  "%v"`+"\n", module.GetName())
    }
}

func (j *Jweb) bootModule(module jwebmodule.Module) {
    name := module.GetName()

    _, ok := j.moduleList[name]
    if ok == true {
        jweberror.Fatal(jweberror.New(`multiple modules registered for name "%v"`, name))
    }

    config := jwebmodule.LoadConfig(module, j.kernel)

    module.Boot(j.kernel, config)

    j.moduleList[name] = module
}

func (j *Jweb) handleError() {
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
