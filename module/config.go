package jwebmodule

import (
    jwebparameter `gitlab.com/drjele-go/jweb/config/parameter`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
    jwebfile `gitlab.com/drjele-go/jweb/utility/file`
)

func LoadConfig(module Module, kernel *jwebkernel.Kernel) *jwebparameter.Yaml {
    var config *jwebparameter.Yaml

    if module.ConfigurationRequired() == false {
        return config
    }

    filePath := kernel.GetRootDir() + `config/` + module.GetName() + `.yaml`

    if jwebfile.Exists(filePath) == false {
        jweberror.Fatal(
            jweberror.New(`the configuration file %v does not exists`, filePath),
        )
    }

    config = jwebparameter.NewYaml([]string{filePath})

    return config
}
