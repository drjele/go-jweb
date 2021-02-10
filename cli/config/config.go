package config

func New(name string, description string) *Config {
    config := Config{
        name:        name,
        description: description,
    }

    return &config
}

type Config struct {
    name        string
    description string
}

func (c *Config) GetName() string {
    return c.name
}

func (c *Config) GetDescription() string {
    return c.description
}
