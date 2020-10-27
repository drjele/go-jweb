package jwebconfig

import (
    `github.com/joho/godotenv`

    jwebconfigcli `gitlab.com/drjele-go/jweb/cli/config`
    jweberror `gitlab.com/drjele-go/jweb/error`
    jwebconfighttp `gitlab.com/drjele-go/jweb/http/config`
    jwebfile `gitlab.com/drjele-go/jweb/utility/file`
)

const (
    EnvDev  = `dev`
    EnvProd = `prod`
)

func New() *Config {
    /** @todo validate minimal params */

    c := Config{}

    c.loadDotEnv()

    c.env = c.GetParam(`ENV`, EnvProd)

    c.http = jwebconfighttp.New(
        c.GetParam(`HTTP_HOST`, `80`),
    )

    c.cli = jwebconfigcli.New(
        c.GetParam(`CLI_NAME`, `jweb`),
        c.GetParam(`CLI_DESCRIPTION`, `Jweb framework`),
    )

    return &c
}

type Config struct {
    params map[string]string
    env    string
    http   *jwebconfighttp.Config
    cli    *jwebconfigcli.Config
}

func (c *Config) GetParam(param string, defaultValue string) string {
    value, ok := c.params[param]

    if ok == false {
        return defaultValue
    }

    return value
}

func (c *Config) GetEnv() string {
    return c.env
}

func (c *Config) GetHttp() *jwebconfighttp.Config {
    return c.http
}

func (c *Config) GetCli() *jwebconfigcli.Config {
    return c.cli
}

func (c *Config) loadDotEnv() {
    var params map[string]string

    files := []string{`.env`}

    if jwebfile.Exists(`.env.local`) {
        files = append(files, `.env.local`)
    }

    params, err := godotenv.Read(files...)
    jweberror.Fatal(err)

    c.params = params
}
