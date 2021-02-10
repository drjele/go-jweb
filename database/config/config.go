package config

import (
    `gitlab.com/drjele-go/jweb/database/connection`
    jweberror `gitlab.com/drjele-go/jweb/error`
)

func New() *Config {
    config := Config{}

    config.connections = connection.Map{}

    return &config
}

type Config struct {
    connections connection.Map
}

func (c *Config) AddConnection(name string, connection *connection.Connection) {
    _, ok := c.connections[name]
    if ok == true {
        jweberror.Fatal(jweberror.New(`duplicate database connection name "%v"`, name))
    }

    c.connections[name] = connection
}

func (c *Config) GetConnection(name string) *connection.Connection {
    _, ok := c.connections[name]
    if ok == false {
        jweberror.Fatal(jweberror.New(`duplicate database connection name "%v"`, name))
    }

    return c.connections[name]
}
