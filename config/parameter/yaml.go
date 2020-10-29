package jwebparameter

import (
    `github.com/knadh/koanf`
    `github.com/knadh/koanf/parsers/yaml`
    `github.com/knadh/koanf/providers/file`

    jweberror `gitlab.com/drjele-go/jweb/error`
)

const (
    PathDelimiter = `.`
)

func NewYaml(files []string) *Yaml {
    pb := Yaml{}

    k := koanf.New(PathDelimiter)

    for _, f := range files {
        if err := k.Load(file.Provider(f), yaml.Parser()); err != nil {
            jweberror.Fatal(
                jweberror.New(`error loading config: %v`, err),
            )
        }
    }

    pb.params = k

    return &pb
}

type Yaml struct {
    params *koanf.Koanf
}

func (y *Yaml) GetParam(param string) interface{} {
    if y.params.Exists(param) == false {
        jweberror.Fatal(
            jweberror.New(`could not find param "%v"`, param),
        )
    }

    return y.params.Get(param)
}

func (y *Yaml) GetParamWithDefault(param string, defaultValue interface{}) interface{} {
    if y.params.Exists(param) == false {
        return defaultValue
    }

    return y.params.Get(param)
}
