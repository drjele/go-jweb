package jwebdatabase

import (
    `fmt`
    `os`

    jwebparameter `gitlab.com/drjele-go/jweb/config/parameter`
    jwebconfig `gitlab.com/drjele-go/jweb/database/config`
    jwebconnection `gitlab.com/drjele-go/jweb/database/connection`
    jwebregistry `gitlab.com/drjele-go/jweb/database/registry`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
    jwebconvert `gitlab.com/drjele-go/jweb/utility/convert`
    jwebmap `gitlab.com/drjele-go/jweb/utility/map`
)

const (
    Name = `database`
)

func New() *Database {
    return &Database{}
}

type Database struct {
    connections jwebconnection.Map
    registry    *jwebregistry.Registry
}

func (d *Database) GetName() string {
    return Name
}

func (d *Database) ConfigurationRequired() bool {
    return true
}

func (d *Database) Boot(kernel *jwebkernel.Kernel, yamlConfig *jwebparameter.Yaml) {
    _ = d.buildConfig(yamlConfig)

    os.Exit(1)

    d.registry = jwebregistry.New()
}

func (d *Database) buildConfig(yamlConfig *jwebparameter.Yaml) *jwebconfig.Config {
    /** @todo add connection names in errors */

    config := jwebconfig.Config{}

    connections, err := jwebconvert.InterfaceToMap(yamlConfig.GetParam(`connections`))
    jweberror.Fatal(err)

    connectionKeys := []string{`driver`, `hostname`, `port`, `username`, `password`, `database`}
    for connectionName, connectionData := range connections {
        connectionDataMap, err := jwebconvert.InterfaceToMap(connectionData)
        jweberror.Fatal(err)

        err = jwebmap.CheckKeysMatch(connectionKeys, connectionDataMap)
        jweberror.Fatal(err)

        fmt.Println(connectionName, connectionData, connectionDataMap)
        os.Exit(12)

        connectionDataMapString, err := jwebconvert.InterfaceToMapString(connectionData)
        jweberror.Fatal(err)

        connection := jwebconnection.New(
            connectionDataMapString[`driver`],
            connectionDataMapString[`hostname`],
            connectionDataMapString[`port`],
            connectionDataMapString[`username`],
            connectionDataMapString[`password`],
            connectionDataMapString[`database`],
        )

        config.AddConnection(connectionName, connection)
    }

    return &config
}
