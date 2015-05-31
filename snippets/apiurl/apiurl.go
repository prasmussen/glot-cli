package apiurl

import (
    "strings"
)

func List(baseUrl string) string {
    return join(baseUrl, "snippets")
}

func Get(baseUrl, id string) string {
    return join(List(baseUrl), id)
}

func join(base string, components ...string) string {
    if !strings.HasSuffix(base, "/") {
        base = base + "/"
    }

    return base + strings.Join(components, "/")
}
