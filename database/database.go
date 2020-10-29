package jwebdatabase

import (
    `fmt`
    `os`

    jwebparameter `gitlab.com/drjele-go/jweb/config/parameter`
    jwebconnection `gitlab.com/drjele-go/jweb/database/connection`
    jwebregistry `gitlab.com/drjele-go/jweb/database/registry`
    jwebkernel `gitlab.com/drjele-go/jweb/kernel`
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

func (d *Database) Validate(config *jwebparameter.Yaml) (err error) {
    connections := config.GetParam(`connections`)

    fmt.Println(connections)
    os.Exit(1)

    return
}

func (d *Database) Boot(kernel *jwebkernel.Kernel, config *jwebparameter.Yaml) {
    d.registry = jwebregistry.New()
}
