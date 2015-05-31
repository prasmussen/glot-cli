package config


func ReadConfig(path string) *Config {
    return &Config{}
}

type Config struct {
    //data map[string]string
    //Run struct {
    //    ApiBaseUrl string `json:"apiBaseUrl"`
    //} `json:"run"`

    //Snippets struct {
    //    ApiBaseUrl string `json:"apiBaseUrl"`
    //} `json:"snippets"`
}

func (self *Config) RunApiBaseUrl() string {
    return "https://run.glot.io/"
}

func (self *Config) RunApiToken() string {
    return "TODO"
}

func (self *Config) SnippetsApiBaseUrl() string {
    return "https://snippets.glot.io/"
}

func (self *Config) SnippetsApiToken() string {
    return "TODO"
}
