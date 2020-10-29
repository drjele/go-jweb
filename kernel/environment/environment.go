package jwebenvironment

import (
    `github.com/joho/godotenv`

    jwebparameter `gitlab.com/drjele-go/jweb/config/parameter`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebfile `gitlab.com/drjele-go/jweb/utility/file`
)

const (
    ModeHttp = `http`
    ModeCli  = `cli`

    EnvDev  = `dev`
    EnvProd = `prod`
)

func New() *Environment {
    e := Environment{}

    /** @todo validate minimal environment vars */
    params := e.loadDotEnv()

    e.params = jwebparameter.NewMap(params)

    return &e
}

type Environment struct {
    params *jwebparameter.Map
}

func (e *Environment) GetDefaultMode() string {
    return e.params.GetParamWithDefault(`DEFAULT_MODE`, ModeHttp)
}

func (e *Environment) GetEnv() string {
    return e.params.GetParamWithDefault(`ENV`, EnvProd)
}

func (e *Environment) GetParam(param string) string {
    return e.params.GetParam(param)
}

func (e *Environment) GetParamWithDefault(param string, defaultValue string) string {
    return e.params.GetParamWithDefault(param, defaultValue)
}

func (e *Environment) loadDotEnv() map[string]string {
    var params map[string]string

    files := []string{`.env`}

    if jwebfile.Exists(`.env.local`) {
        files = append(files, `.env.local`)
    }

    params, err := godotenv.Read(files...)
    jweberror.Fatal(err)

    return params
}
