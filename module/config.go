package module

import (
    `log`
    `strings`

    `gitlab.com/drjele-go/jweb/config/parameter`
    jweberror `gitlab.com/drjele-go/jweb/error`
    `gitlab.com/drjele-go/jweb/kernel`
    `gitlab.com/drjele-go/jweb/kernel/environment`
    `gitlab.com/drjele-go/jweb/utility/file`
)

func LoadConfig(module Module, kernel *kernel.Kernel) *parameter.Yaml {
    var config *parameter.Yaml

    if module.ConfigurationRequired() == false {
        return config
    }

    filePath := kernel.GetRootDir() + `config/` + module.GetName() + `.yaml`
    log.Printf(`loading config "%v"`+"\n", filePath)

    if file.Exists(filePath) == false {
        jweberror.Fatal(
            jweberror.New(`the configuration file %v does not exists`, filePath),
        )
    }

    config = parameter.NewYamlFromFiles([]string{filePath})

    config = parseConfig(config, kernel.GetEnvironment())

    return config
}

func parseConfig(config *parameter.Yaml, environment *environment.Environment) *parameter.Yaml {
    m := map[string]interface{}{}

    for _, key := range config.Keys() {
        param := config.GetParam(key)

        stringParam, ok := param.(string)
        if ok == false {
            /** @todo maybe also log */
            m[key] = param
            continue
        }

        m[key] = parseParam(stringParam, environment)
    }

    return parameter.NewYamlFromMap(m)
}

func parseParam(stringParam string, environment *environment.Environment) string {
    if (strings.HasPrefix(stringParam, `%env(`) && strings.HasSuffix(stringParam, `)%`)) == false {
        return stringParam
    }

    stringParam = stringParam[5:]
    stringParam = stringParam[:len(stringParam)-2]

    if environment.HasParam(stringParam) == false {
        jweberror.Fatal(
            jweberror.New(`the env param "%v" was not found`, stringParam),
        )
    }

    return environment.GetParam(stringParam)
}
