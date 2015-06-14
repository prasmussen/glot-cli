package run

import (
    "fmt"
    "strings"
    "os"
    "path/filepath"
    "io/ioutil"
    "./apiurl"
    "./api"
    "../util"
    "../language"
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
    lang, ok := language.DetermineLanguage(path)
    if !ok {
        fmt.Fprintln(os.Stderr, "Unknown filetype")
        return
    }

    files, err := readFiles(path, lang)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read files: %s\n", err.Error())
        return
    }

    url := apiurl.Run(cfg.RunApiBaseUrl(), lang, version)
    token := cfg.RunApiToken()
    runResult, err := api.Run(url, token, files)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to run code: %s\n", err.Error())
        return
    }

    fmt.Fprint(os.Stdout, runResult.Stdout)
    fmt.Fprint(os.Stderr, runResult.Stderr)
    fmt.Fprint(os.Stderr, runResult.Error)
}

func readFiles(path, lang string) ([]*api.File, error) {
    isRegular, err := util.IsRegularFile(path)
    if err != nil {
        return nil, err
    } else if !isRegular {
        return nil, fmt.Errorf("File (%s) is not a regular file", path)
    }
        
    root := filepath.Dir(path)

    paths, err := collectPaths(root, lang)
    if err != nil {
        return nil, err
    }

    // Ensure that the given file is first in the paths slice
    paths = ensureFirst(paths, path)

    files, err := pathsToFiles(paths)
    if err != nil {
        return nil, err
    }

    return files, err
}

func ensureFirst(paths []string, path string) []string {
    var newPaths = make([]string, 0)

    for _, p := range paths {
        if p == path {
            newPaths = append(newPaths, p)
        }
    }

    for _, p := range paths {
        if p != path {
            newPaths = append(newPaths, p)
        }
    }

    return newPaths
}

func collectPaths(root, lang string) ([]string, error) {
    var paths = make([]string, 0)

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip hidden files / dirs
        if strings.HasPrefix(path, ".") {
            return nil
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        // Skip non allowed files
        ext := filepath.Ext(path)[1:]
        if !language.AllowedExtension(ext, lang) {
            return nil
        }

        paths = append(paths, path)

        return nil
    })

    return paths, err
}

func pathsToFiles(paths []string) ([]*api.File, error) {
    var files = make([]*api.File, 0)

    for _, path := range paths {
        // Read file content
        bytes, err := ioutil.ReadFile(path)
        if err != nil {
            return nil, err
        }

        files = append(files, &api.File{
            Name: path,
            Content: string(bytes),
        })
    }

    return files, nil
}
