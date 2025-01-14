package connection

import (
    jweberror `gitlab.com/drjele-go/jweb/error`
    `gitlab.com/drjele-go/jweb/utility`
)

const (
    DriverMysql = `mysql`
    DriverMongo = `mongo`
)

type Map map[string]*Connection

func New(driver, hostname, port, username, password, database string) *Connection {
    if utility.StringInSlice(driver, []string{DriverMysql, DriverMongo}) == false {
        jweberror.Fatal(jweberror.New(`invalid driver "%v"`, driver))
    }

    return &Connection{
        driver:   driver,
        hostname: hostname,
        port:     port,
        username: username,
        password: password,
        database: database,
    }
}

type Connection struct {
    driver   string
    hostname string
    port     string
    username string
    password string
    database string
}

func (c *Connection) GetDriver() string {
    return c.driver
}

func (c *Connection) GetHostname() string {
    return c.hostname
}

func (c *Connection) GetPort() string {
    return c.port
}

func (c *Connection) GetUsername() string {
    return c.username
}

func (c *Connection) GetPassword() string {
    return c.password
}

func (c *Connection) GetDatabase() string {
    return c.database
}
