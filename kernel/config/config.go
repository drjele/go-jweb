package config

import (
    configcli `gitlab.com/drjele-go/jweb/cli/config`
    confighttp `gitlab.com/drjele-go/jweb/http/config`
    `gitlab.com/drjele-go/jweb/kernel/environment`
)

func New(environment *environment.Environment) *Config {
    c := Config{}

    c.http = confighttp.New(
        environment.GetParam(`HTTP_HOST`),
    )

    c.cli = configcli.New(
        environment.GetParam(`CLI_NAME`),
        environment.GetParam(`CLI_DESCRIPTION`),
    )

    return &c
}

type Config struct {
    http *confighttp.Config
    cli  *configcli.Config
}

func (c *Config) GetHttp() *confighttp.Config {
    return c.http
}

func (c *Config) GetCli() *configcli.Config {
    return c.cli
}
