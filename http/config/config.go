package jwebconfig

func New(host string) *Config {
    return &Config{
        host: host,
    }
}

type Config struct {
    host string
}

func (c *Config) GetHost() string {
    return c.host
}
