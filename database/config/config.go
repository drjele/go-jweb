package jwebconfig

import (
    jwebconnection `gitlab.com/drjele-go/jweb/database/connection`
    jweberror `gitlab.com/drjele-go/jweb/error`
)

type Config struct {
    connections jwebconnection.Map
}

func (c *Config) AddConnection(name string, connection *jwebconnection.Connection) {
    _, ok := c.connections[name]
    if ok == true {
        jweberror.Fatal(jweberror.New(`duplicate database connection name "%v"`, name))
    }

    c.connections[name] = connection
}

func (c *Config) GetConnection(name string) *jwebconnection.Connection {
    _, ok := c.connections[name]
    if ok == false {
        jweberror.Fatal(jweberror.New(`duplicate database connection name "%v"`, name))
    }

    return c.connections[name]
}
