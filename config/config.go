package config

import (
    "os"
    "encoding/json"
)

func ReadConfig(path string) (*Config, error) {
    cfg := &Config{}
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    err = json.NewDecoder(f).Decode(cfg)
    return cfg, err
}

type Config struct {
    Run struct {
        BaseUrl string `json:"baseUrl"`
        Token string `json:"token"`
    } `json:"run"`

    Snippets struct {
        BaseUrl string `json:"baseUrl"`
        Token string `json:"token"`
    } `json:"snippets"`
}

func (self *Config) RunApiBaseUrl() string {
    return self.Run.BaseUrl
}

func (self *Config) RunApiToken() string {
    return self.Snippets.Token
}

func (self *Config) SnippetsApiBaseUrl() string {
    return self.Snippets.BaseUrl
}

func (self *Config) SnippetsApiToken() string {
    return self.Snippets.Token
}
