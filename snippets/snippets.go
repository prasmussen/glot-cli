package snippets

import (
    "fmt"
    "os"
    "strings"
    "strconv"
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

    printMeta(snippet)
}

func findMainFile(names []string) string {
    for _, name := range names {
        base := filepath.Base(name)
        if strings.HasPrefix(strings.ToLower(base), "main") {
            return base
        }
    }

    return names[0]
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

func printMeta(snippet *api.Snippet) {
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
