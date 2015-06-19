package snippets

import (
    "fmt"
    "strings"
    "strconv"
    "os"
    "path/filepath"
    "io/ioutil"
    "../language"
    "../util"
    "./apiurl"
    "./api"
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

func ListSnippets(cfg config) {
    url := apiurl.List(cfg.SnippetsApiBaseUrl())
    token := cfg.SnippetsApiToken()
    snippets, err := api.List(url, token)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to list snippets: %s\n", err.Error())
        return
    }

    // Define column format
    format := "%-22s %-12s %-12s %s\n"

    // Print header
    fmt.Printf(format, "Created", "Id", "Language", "Title")

    // Print snippets
    for _, snippet := range snippets {
        fmt.Printf(format, snippet.Created, snippet.Id, snippet.Language, snippet.Title)
    }
}

func PrintMeta(cfg config, id string) {
    url := apiurl.Get(cfg.SnippetsApiBaseUrl(), id)
    snippet, err := api.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to get snippets: %s\n", err.Error())
        return
    }

    fmt.Printf("Id: %s\n", snippet.Id)
    fmt.Printf("Title: %s\n", snippet.Title)
    fmt.Printf("Language: %s\n", snippet.Language)
    fmt.Printf("Public: %s\n", strings.Title(strconv.FormatBool(snippet.Public)))
    fmt.Printf("Created: %s\n", snippet.Created)
    fmt.Printf("Modified: %s\n", snippet.Modified)
    fmt.Printf("Owner: %s\n", snippet.Owner)
    fmt.Printf("Files hash: %s\n", snippet.FilesHash)
    fmt.Printf("Web Url: %s\n", strings.Replace(snippet.Url, "snippets.glot.io", "glot.io", 1))
    fmt.Printf("Api Url: %s\n", snippet.Url)
}

func Clone(cfg config, id string) {
    basePath := id
    fmt.Printf("Cloning into '%s'...\n", basePath)

    // Check if destination dir already exists
    if util.FileExists(basePath) {
        fmt.Fprintf(os.Stderr, "Destination path '%s' already exists\n", basePath)
        return
    }

    // Get snippet from api
    url := apiurl.Get(cfg.SnippetsApiBaseUrl(), id)
    snippet, err := api.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to get snippets: %s\n", err.Error())
        return
    }

    // Write snippet to disk
    writeSnippet(basePath, snippet)
}

func Pull(cfg config) {
    fmt.Printf("Pulling...\n")

    // Read meta file
    meta, err := findReadMetaFile()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        return
    }

    // Get snippet from api
    url := apiurl.Get(cfg.SnippetsApiBaseUrl(), meta.Id)
    snippet, err := api.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to get snippets: %s\n", err.Error())
        return
    }

    // Write snippet to disk
    writeSnippet(".", snippet)
}

func Push(cfg config) {
    fmt.Printf("Pushing...\n")

    basePath, ok := findGlotPath()
    if !ok {
        pushNew(cfg, basePath)
    } else {
        pushUpdate(cfg, basePath)
    }
}

func Delete(cfg config, id string) {
    fmt.Println("Deleting...")

    url := apiurl.Delete(cfg.SnippetsApiBaseUrl(), id)
    token := cfg.SnippetsApiToken()
    err := api.Delete(url, token)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to delete snippet: %s\n", err.Error())
        return
    }

    fmt.Printf("Done\n")
}


func pushNew(cfg config, basePath string) {
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

    writeSnippet(".", snippet)
    fmt.Printf("Saved snippet %s\n", snippet.Id)
}

func pushUpdate(cfg config, basePath string) {
    fmt.Println("Updating...")
    // Read meta file
    meta, err := findReadMetaFile()
    if err != nil {

    }

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

    // TODO: Preserve main file
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

    url := apiurl.Update(cfg.SnippetsApiBaseUrl(), meta.Id)
    token := cfg.SnippetsApiToken()
    snippet, err := api.Update(url, token, snippetData)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to update snippet: %s\n", err.Error())
        return
    }

    writeSnippet(".", snippet)
    fmt.Printf("Updated snippet %s\n", snippet.Id)
}
