package jwebmodule

import (
    `log`
    `strings`

    jwebparameter `gitlab.com/drjele-go/jweb/config/parameter`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
    jwebenvironment `gitlab.com/drjele-go/jweb/kernel/environment`
    jwebfile `gitlab.com/drjele-go/jweb/utility/file`
)

func LoadConfig(module Module, kernel *jwebkernel.Kernel) *jwebparameter.Yaml {
    var config *jwebparameter.Yaml

    if module.ConfigurationRequired() == false {
        return config
    }

    filePath := kernel.GetRootDir() + `config/` + module.GetName() + `.yaml`
    log.Printf(`loading config "%v"`+"\n", filePath)

    if jwebfile.Exists(filePath) == false {
        jweberror.Fatal(
            jweberror.New(`the configuration file %v does not exists`, filePath),
        )
    }

    config = jwebparameter.NewYamlFromFiles([]string{filePath})

    config = parseConfig(config, kernel.GetEnvironment())

    return config
}

func parseConfig(config *jwebparameter.Yaml, environment *jwebenvironment.Environment) *jwebparameter.Yaml {
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

    return jwebparameter.NewYamlFromMap(m)
}

func parseParam(stringParam string, environment *jwebenvironment.Environment) string {
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
