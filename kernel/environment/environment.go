package environment

import (
    `github.com/joho/godotenv`

    `gitlab.com/drjele-go/jweb/config/parameter`
    jweberror `gitlab.com/drjele-go/jweb/error`
    `gitlab.com/drjele-go/jweb/utility`
)

const (
    ModeHttp = `http`
    ModeCli  = `cli`

    EnvDev  = `dev`
    EnvProd = `prod`
)

func New(rootDir string) *Environment {
    environment := Environment{}

    /** @todo validate minimal environment vars */
    params := environment.loadDotEnv(rootDir)

    fullMap := parameter.NewMap(params)
    /** parameters with defaults should be initialized here */
    environment.defaultMode = fullMap.GetParamWithDefault(`DEFAULT_MODE`, ModeHttp)
    delete(params, `DEFAULT_MODE`)
    environment.env = fullMap.GetParamWithDefault(`ENV`, EnvProd)
    delete(params, `ENV`)

    environment.params = parameter.NewMap(params)

    return &environment
}

type Environment struct {
    params      *parameter.Map
    defaultMode string
    env         string
}

func (e *Environment) GetDefaultMode() string {
    return e.defaultMode
}

func (e *Environment) GetEnv() string {
    return e.env
}

func (e *Environment) GetParam(param string) string {
    return e.params.GetParam(param)
}

func (e *Environment) GetParamWithDefault(param string, defaultValue string) string {
    return e.params.GetParamWithDefault(param, defaultValue)
}

func (e *Environment) HasParam(param string) bool {
    return e.params.HasParam(param)
}

func (e *Environment) loadDotEnv(rootDir string) map[string]string {
    var params map[string]string

    files := []string{rootDir + `.env`}

    envLocal := rootDir + `.env.local`
    if utility.Exists(envLocal) {
        files = append(files, envLocal)
    }

    params, err := godotenv.Read(files...)
    jweberror.Fatal(err)

    return params
}
