package run

import (
    "fmt"
    "os"
    "../util"
    "../language"
    "github.com/prasmussen/glot-api-lib/go/run"
    apiurl "github.com/prasmussen/glot-api-lib/go/run/url"
)

type config interface {
    RunApiBaseUrl() string
    RunApiToken() string
}

func ListLanguages(cfg config) {
    url := apiurl.ListLanguages(cfg.RunApiBaseUrl())
    languages, err := api.ListLanguages(url)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to list languages: %s\n", err.Error())
        return
    }

    for _, lang := range languages {
        fmt.Println(lang.Name)
    }
}

func ListVersions(cfg config, lang string) {
    url := apiurl.ListVersions(cfg.RunApiBaseUrl(), lang)
    versions, err := api.ListVersions(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to list versions: %s\n", err.Error())
        return
    }

    for _, vsn := range versions {
        fmt.Println(vsn.Version)
    }
}

func Run(cfg config, version, path string) {
    files, err := util.ReadFile(path)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read file: %s\n", err.Error())
        return
    }

    lang, ok := language.DetermineLanguage(path)
    if !ok {
        fmt.Fprintln(os.Stderr, "Unknown filetype")
        return
    }

    url := apiurl.Run(cfg.RunApiBaseUrl(), lang, version)
    token := cfg.RunApiToken()
    runResult, err := api.Run(url, token, toApiFiles(files))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to run code: %s\n", err.Error())
        return
    }

    fmt.Fprint(os.Stdout, runResult.Stdout)
    fmt.Fprint(os.Stderr, runResult.Stderr)
    fmt.Fprint(os.Stderr, runResult.Error)
}

func toApiFiles(files []*util.File) []*api.File {
    apiFiles := make([]*api.File, 0, len(files))
    for _, f := range files {
        apiFiles = append(apiFiles, &api.File{
            Name: f.Name,
            Content: f.Content,
        })
    }
    return apiFiles
}
