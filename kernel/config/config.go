package jwebconfig

import (
    jwebconfigcli `gitlab.com/drjele-go/jweb/cli/config`
    jwebconfighttp `gitlab.com/drjele-go/jweb/http/config`
    jwebenvironment `gitlab.com/drjele-go/jweb/kernel/environment`
)

func New(environment *jwebenvironment.Environment) *Config {
    c := Config{}

    c.http = jwebconfighttp.New(
        environment.GetParam(`HTTP_HOST`),
    )

    c.cli = jwebconfigcli.New(
        environment.GetParam(`CLI_NAME`),
        environment.GetParam(`CLI_DESCRIPTION`),
    )

    return &c
}

type Config struct {
    http *jwebconfighttp.Config
    cli  *jwebconfigcli.Config
}

func (c *Config) GetHttp() *jwebconfighttp.Config {
    return c.http
}

func (c *Config) GetCli() *jwebconfigcli.Config {
    return c.cli
}
