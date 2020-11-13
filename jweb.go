package jweb

import (
    `flag`
    `fmt`
    `log`
    `os`
    `runtime/debug`
    `sync`

    jwebcli `gitlab.com/drjele-go/jweb/src/cli`
    command `gitlab.com/drjele-go/jweb/src/cli/command`
    jweberror `gitlab.com/drjele-go/jweb/src/error`
    jwebhttp `gitlab.com/drjele-go/jweb/src/http`
    route `gitlab.com/drjele-go/jweb/src/http/routing/route`
    jwebkernel `gitlab.com/drjele-go/jweb/src/kernel`
    environment `gitlab.com/drjele-go/jweb/src/kernel/environment`
    jwebmodule `gitlab.com/drjele-go/jweb/src/module`
)

var (
    jweb *Jweb
    once sync.Once
)

func Get(
    moduleList jwebmodule.List,
) *Jweb {
    once.Do(
        func() {
            jweb = newInstance(moduleList)
        },
    )

    return jweb
}

func newInstance(moduleList jwebmodule.List) *Jweb {
    defer handleError()

    jweb := Jweb{
        routeList:   route.List{},
        commandList: command.List{},
    }

    dir, err := os.Getwd()
    jweberror.Fatal(err)

    /** @todo maybe not add the trailing slash */
    kernel := jwebkernel.New(dir + `/`)

    jweb.kernel = kernel
    jweb.errorHandler = jweberror.NewHandler(
        kernel.GetEnvironment().GetEnv(),
    )

    jweb.bootModules(moduleList)

    return &jweb
}

type Jweb struct {
    kernel       *jwebkernel.Kernel
    errorHandler *jweberror.Handler
    moduleList   jwebmodule.Map
    routeList    route.List
    commandList  command.List
}

func (j *Jweb) SetRouteList(routeList route.List) {
    j.routeList = routeList
}

func (j *Jweb) SetCommandList(commandList command.List) {
    j.commandList = commandList
}

func (j *Jweb) Run() {
    defer j.handleError()

    switch j.kernel.GetFlags().GetMode() {
    case environment.ModeHttp:
        jwebhttp.Run(
            j.kernel.GetEnvironment(),
            j.kernel.GetConfig(),
            j.routeList,
        )
        break
    case environment.ModeCli:
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

func (j *Jweb) Panic(err error) {
    if err == nil {
        return
    }

    if j.errorHandler == nil {
        panic(err)

        return
    }

    j.errorHandler.Panic(err)
}

func (j *Jweb) Fatal(err error) {
    if err == nil {
        return
    }

    if j.errorHandler == nil {
        defer func() {
            debug.PrintStack()
        }()

        log.Fatal(err)

        return
    }

    j.errorHandler.Fatal(err)
}

func (j *Jweb) bootModules(moduleList jwebmodule.List) {
    j.moduleList = jwebmodule.Map{}

    for _, module := range moduleList {
        log.Printf(`boot start "%v"`+"\n", module.GetName())

        j.bootModule(module)

        log.Printf(`boot end "%v"`+"\n", module.GetName())
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
    handleError()
}

func handleError() {
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
