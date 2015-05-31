package api

import (
    "io"
    "net/http"
    "encoding/json"
)

type File struct {
    Name string `json:"name"`
    Content string `json:"content"`
}

type RunResult struct {
    Stdout string `json:"stdout"`
    Stderr string `json:"stderr"`
    Error string `json:"error"`
}

type Language struct {
    Name string `json:"name"`
    Url string `json:"url"`
}

type Version struct {
    Version string `json:"version"`
    Url string `json:"url"`
}

func ListLanguages(url string) ([]*Language, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    languages := make([]*Language, 0)
    err = json.NewDecoder(res.Body).Decode(&languages)
    if err != nil {
        return nil, err
    }

    return languages, err
}

func ListVersions(url string) ([]*Version, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    versions := make([]*Version, 0)
    err = json.NewDecoder(res.Body).Decode(&versions)
    if err != nil {
        return nil, err
    }

    return versions, err
}

func Run(url, token string, files []*File) (*RunResult, error) {
    jsonReader, jsonWriter := io.Pipe() 
    defer jsonReader.Close()

    // Start async writer
    go func() {
        json.NewEncoder(jsonWriter).Encode(files)
        jsonWriter.Close()
    }()

    req, err := http.NewRequest("POST", url, jsonReader)
    req.Header.Set("Authorization", "Token " + token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    runResult := &RunResult{}
    err = json.NewDecoder(res.Body).Decode(&runResult)
    if err != nil {
        return nil, err
    }

    return runResult, err
}
