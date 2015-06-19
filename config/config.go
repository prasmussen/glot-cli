package config

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "../util"
)

const (
    defaultRunBaseUrl = "https://run.glot.io"
    defaultSnippetsBaseUrl = "https://snippets.glot.io"
)


type Config struct {
    Run runConfig `json:"run"`
    Snippets snippetsConfig`json:"snippets"`
}

type runConfig struct {
    BaseUrl string `json:"baseUrl"`
    Token string `json:"token"`
}

type snippetsConfig struct {
    BaseUrl string `json:"baseUrl"`
    Token string `json:"token"`
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

func DefaultConfig() *Config {
    return &Config{
        Run: runConfig{
            BaseUrl: defaultRunBaseUrl,
            Token: "",
        },
        Snippets: snippetsConfig{
            BaseUrl: defaultSnippetsBaseUrl,
            Token: "",
        },
    }
}

func ReadConfig(path string) (*Config, error) {
    cfg := &Config{}
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    err = json.NewDecoder(f).Decode(cfg)
    return cfg, err
}

func SaveConfig(path string, cfg *Config) error {
    data, err := json.MarshalIndent(cfg, "", "    ")
    if err != nil {
        return err
    }

    if err = util.Mkdir(path); err != nil {
        return err
    }
    return ioutil.WriteFile(path, data, 0600)
}
