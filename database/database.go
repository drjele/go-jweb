package database

import (
    `gitlab.com/drjele-go/jweb/config/parameter`
    jwebconfig `gitlab.com/drjele-go/jweb/database/config`
    jwebconnection `gitlab.com/drjele-go/jweb/database/connection`
    jwebmanager `gitlab.com/drjele-go/jweb/database/manager`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
    `gitlab.com/drjele-go/jweb/utility`
)

const (
    Name = `database`
)

func New() *Database {
    database := Database{}

    database.managers = jwebmanager.Map{}

    return &database
}

type Database struct {
    kernel   *jwebkernel.Kernel
    config   *jwebconfig.Config
    managers jwebmanager.Map
}

func (d *Database) GetName() string {
    return Name
}

func (d *Database) ConfigurationRequired() bool {
    return true
}

func (d *Database) Boot(kernel *jwebkernel.Kernel, yamlConfig *parameter.Yaml) {
    d.kernel = kernel

    d.config = d.buildConfig(yamlConfig)
}

func (d *Database) GetManager(name string) jwebmanager.Manager {
    _, ok := d.managers[name]

    if ok == false {
        d.managers[name] = d.initManager(d.config.GetConnection(name))
    }

    manager, _ := d.managers[name]

    return manager
}

func (d *Database) initManager(connection *jwebconnection.Connection) jwebmanager.Manager {
    var db jwebmanager.Manager

    switch connection.GetDriver() {
    case jwebconnection.DriverMysql:
        db = jwebmanager.NewMysql(connection, d.kernel.GetEnvironment().GetEnv())
        break
    case jwebconnection.DriverMongo:
        db = jwebmanager.NewMongo(connection)
        break
    default:
        jweberror.Fatal(jweberror.New(`invalid connection driver "%v"`, connection.GetDriver()))
    }

    return db
}

func (d *Database) buildConfig(yamlConfig *parameter.Yaml) *jwebconfig.Config {
    /** @todo add connection names in errors */

    config := jwebconfig.New()

    connections, err := utility.InterfaceToMap(yamlConfig.GetParam(`connections`))
    jweberror.Fatal(err)

    connectionKeys := []string{`driver`, `hostname`, `port`, `username`, `password`, `database`}
    for connectionName, connectionData := range connections {
        connectionDataMap, err := utility.InterfaceToMap(connectionData)
        jweberror.Fatal(err)

        err = utility.CheckKeysMatch(connectionKeys, connectionDataMap)
        jweberror.Fatal(err)

        connectionDataMapString, err := utility.MapInterfaceToString(connectionDataMap)
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

    return config
}
