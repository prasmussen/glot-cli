package util

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
    "path/filepath"
)

func FileExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil
}

func IsRegularFile(path string) (bool, error) {
    f, err := os.Open(path)
    if err != nil {
        return false, err
    }

    info, err := f.Stat()
    if err != nil {
        return false, err
    }

    return info.Mode().IsRegular(), nil
}

type File struct {
    Name string
    Content string
}

func ReadFiles(path string) ([]*File, error) {
    isRegular, err := IsRegularFile(path)
    if err != nil {
        return nil, err
    } else if !isRegular {
        return nil, fmt.Errorf("File (%s) is not a regular file", path)
    }

    root := filepath.Dir(path)

    paths, err := collectPaths(root)
    if err != nil {
        return nil, err
    }

    // Ensure that the given file is first in the paths list
    paths = ensureFirst(path, paths)

    files, err := pathsToFiles(paths)
    if err != nil {
        return nil, err
    }

    return files, err
}

func collectPaths(root string) ([]string, error) {
    var paths = make([]string, 0)

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip directories
        if info.IsDir() {
            return nil
        }

        // Skip hidden files
        if strings.HasPrefix(info.Name(), ".") {
            return nil
        }

        paths = append(paths, path)

        return nil
    })

    return paths, err
}

func pathsToFiles(paths []string) ([]*File, error) {
    var files = make([]*File, 0)

    for _, path := range paths {
        // Read file content
        bytes, err := ioutil.ReadFile(path)
        if err != nil {
            return nil, err
        }

        files = append(files, &File{
            Name: path,
            Content: string(bytes),
        })
    }

    return files, nil
}

func ensureFirst(firstItem string, items []string) []string {
    newItems := []string{firstItem}

    // Append the rest of the items
    for _, item := range items {
        if item != firstItem {
            newItems = append(newItems, item)
        }
    }

    return newItems
}
