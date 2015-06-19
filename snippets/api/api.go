package api

import (
    "fmt"
    "io"
    "io/ioutil"
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
        json.NewEncoder(jsonWriter).Encode(data)
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

func Update(url, token string, data *SnippetData) (*Snippet, error) {
    jsonReader, jsonWriter := io.Pipe() 
    defer jsonReader.Close()

    // Start async writer
    go func() {
        json.NewEncoder(jsonWriter).Encode(data)
        jsonWriter.Close()
    }()

    req, err := http.NewRequest("PUT", url, jsonReader)
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

func Delete(url, token string) error {
    req, err := http.NewRequest("DELETE", url, nil)
    req.Header.Set("Authorization", "Token " + token)

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return err
    }

    if res.StatusCode == 204 {
        return nil
    } else if res.StatusCode == 403 {
        return fmt.Errorf("Not allowed to delete this snippet")
    } else if res.StatusCode == 404 {
        return fmt.Errorf("Snippet not found")
    }

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)
    msg := fmt.Sprintf("Unexpected status code: %d", res.StatusCode)
    if len(body) > 0 {
        msg = fmt.Sprintf("%s\n%s", msg, string(body))
    }
    return fmt.Errorf(msg)
}
