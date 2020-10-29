package jwebparameter

import (
    jweberror `gitlab.com/drjele-go/jweb/error`
)

func NewMap(params map[string]string) *Map {
    return &Map{params: params}
}

type Map struct {
    params map[string]string
}

func (m *Map) GetParam(param string) string {
    value, ok := m.params[param]

    if ok == false {
        jweberror.Fatal(
            jweberror.New(`could not find param "%v"`, param),
        )
    }

    return value
}

func (m *Map) GetParamWithDefault(param string, defaultValue string) string {
    value, ok := m.params[param]

    if ok == false {
        return defaultValue
    }

    return value
}
