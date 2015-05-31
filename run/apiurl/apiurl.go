package apiurl

import (
    "strings"
)

func ListLanguages(baseUrl string) string {
    return join(baseUrl, "languages")
}

func ListVersions(baseUrl, lang string) string {
    return join(ListLanguages(baseUrl), lang)
}

func Run(baseUrl, lang, version string) string {
    return join(ListVersions(baseUrl, lang), version)
}

func join(base string, components ...string) string {
    if !strings.HasSuffix(base, "/") {
        base = base + "/"
    }

    return base + strings.Join(components, "/")
}
