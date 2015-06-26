package snippets

import (
    "fmt"
    "os"
    "path/filepath"
    "io/ioutil"
    "../language"
    "../util"
    "github.com/prasmussen/glot-api-lib/go/snippets"
    apiurl "github.com/prasmussen/glot-api-lib/go/snippets/url"
)

type config interface {
    SnippetsApiBaseUrl() string
    SnippetsApiToken() string
}

func NewSnippet(lang string) {
    content, ok := language.DefaultContent(lang)
    if !ok {
        fmt.Fprintln(os.Stderr, "Language not supported")
        return
    }

    fname, _ := language.DefaultFname(lang)
    dir := "untitled"
    path := filepath.Join(dir, fname)

    err := os.Mkdir(dir, 0775)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create directory: %s\n", err.Error())
        return
    }

    err = ioutil.WriteFile(path, []byte(content), 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to write file: %s\n", err.Error())
        return
    }

    fmt.Printf("Created %s\n", path)
}

func Publish(cfg config) {
    fmt.Printf("Publishing...\n")
    basePath, _ := findGlotPath()

    f, err := os.Open(".")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read files: %s\n", err.Error())
        return
    }

    names, err := f.Readdirnames(0)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read files: %s\n", err.Error())
        return
    }

    mainFile := findMainFile(names)

    files, err := util.ReadFiles(mainFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to read files: %s\n", err.Error())
        return
    }

    lang, ok := language.DetermineLanguage(mainFile)
    if !ok {
        fmt.Fprintln(os.Stderr, "Failed to determine language")
        return
    }

    snippetData := &api.SnippetData{
        Language: lang,
        Title: "untitled",
        Public: false,
        Files: toApiFiles(files),
    }

    url := apiurl.Create(cfg.SnippetsApiBaseUrl())
    token := cfg.SnippetsApiToken()
    snippet, err := api.Create(url, token, snippetData)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create snippet: %s\n", err.Error())
        return
    }

    // Write meta file to disk
    metaFilePath := filepath.Join(basePath, ".glot", "meta")
    err = writeMetaFile(metaFilePath, snippet.MetaSnippet)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to write meta file to disk: %s\n", err.Error())
        return
    }

    fmt.Printf("Saved snippet %s\n", snippet.Id)
}
