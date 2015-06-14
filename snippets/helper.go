package snippets

import (
    "os"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "path/filepath"
    "strings"
    "./api"
    "../util"
)


// Write snippet to disk
func writeSnippet(basePath string, snippet *api.Snippet) {
    // Write files to disk
    err := writeFiles(basePath, snippet.Files)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to write file to disk: %s\n", err.Error())
        return
    }

    // Write meta file to disk
    metaFilePath := filepath.Join(basePath, ".glot", "meta")
    err = writeMetaFile(metaFilePath, snippet.MetaSnippet)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to write meta file to disk: %s\n", err.Error())
        return
    }
}

// Writes files to disk
func writeFiles(basePath string, files []*api.File) error {
    for _, file := range files {
        // Get absolute path to file inside basePath
        absPath := filepath.Join(basePath, file.Name)

        err := writeFile(absPath, []byte(file.Content))
        if err != nil {
            return err
        }

        fmt.Printf("Created %s\n", absPath)
    }

    return nil
}

func writeFile(path string, data []byte) error {
    // Create all parent dirs
    err := os.MkdirAll(filepath.Dir(path), 0775)
    if err != nil {
        return err
    }

    // Write file to disk
    return ioutil.WriteFile(path, data, 0664)
}

func writeMetaFile(path string, meta *api.MetaSnippet) error {
    data, err := json.MarshalIndent(meta, "", "  ")
    if err != nil {
        return err
    }
    return writeFile(path, data)
}

func findReadMetaFile() (*api.MetaSnippet, error) {
    // Find .glot dir
    basePath, ok := findGlotPath()
    if !ok {
        return nil, fmt.Errorf(".glot directory not found")
    }

    // Find meta file
    metaFilePath := filepath.Join(basePath, "meta")
    if !util.FileExists(metaFilePath) {
        return nil, fmt.Errorf("meta file not found")
    }

    // Read meta file
    meta, err := readMetaFile(metaFilePath)
    if err != nil {
        return nil, fmt.Errorf("Failed to read meta file: %s", err.Error())
    }

    return meta, nil
}

func findGlotPath() (string, bool) {
    path := ".glot"
    return path, util.FileExists(path)
}

func readMetaFile(path string) (*api.MetaSnippet, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }

    meta := &api.MetaSnippet{}
    err = json.NewDecoder(f).Decode(meta)
    return meta, err
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
