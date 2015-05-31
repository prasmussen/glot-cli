package api

import (
    "net/http"
    "encoding/json"
)

type Snippet struct {
    *MetaSnippet
    Files []*File `json:"files"`
}

type MetaSnippet struct {
    Url string `json:"url"`
    Id string `json:"id"`
    Created string `json:"created"`
    Modified string `json:"modified"`
    Language string `json:"language"`
    Title string `json:"title"`
    Public bool `json:"public"`
    Owner string `json:"owner"`
    FilesHash string `json:"files_hash"`
}

type SnippetData struct {
    Language string `json:"language"`
    Title string `json:"title"`
    Public bool `json:"public"`
    Files []*File `json:"files"`
}

type File struct {
    Name string `json:"name"`
    Content string `json:"content"`
}

func List(url, token string) ([]*MetaSnippet, error) {
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Token " + token)

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    snippets := make([]*MetaSnippet, 0)
    err = json.NewDecoder(res.Body).Decode(&snippets)
    if err != nil {
        return nil, err
    }

    return snippets, err
}

func Get(url string) (*Snippet, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    snippet := &Snippet{}
    err = json.NewDecoder(res.Body).Decode(snippet)
    if err != nil {
        return nil, err
    }

    return snippet, err
}

func Create(url, token string, data *SnippetData) (*Snippet, error) {
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

    snippet := &Snippet{}
    err = json.NewDecoder(res.Body).Decode(&snippet)
    if err != nil {
        return nil, err
    }

    return snippet, err
}
