package language

import (
    "path/filepath"
)

var defaultContent = map[string]string{
    "assembly": assemblyContent,
    "bash": bashContent,
    "c": cContent,
    "clojure": clojureContent,
    "cpp": cppContent,
    "csharp": csharpContent,
    "elixir": elixirContent,
    "erlang": erlangContent,
    "fsharp": fsharpContent,
    "go": goContent,
    "haskell": haskellContent,
    "java": javaContent,
    "javascript": javascriptContent,
    "lua": luaContent,
    "perl": perlContent,
    "php": phpContent,
    "python": pythonContent,
    "ruby": rubyContent,
    "rust": rustContent,
    "scala": scalaContent,
    "plaintext": plaintextContent,
}

var fileExt = map[string]string{
    "assembly": "asm",
    "bash": "sh",
    "c": "c",
    "clojure": "clj",
    "cpp": "cpp",
    "csharp": "cs",
    "elixir": "ex",
    "erlang": "erl",
    "fsharp": "fs",
    "go": "go",
    "haskell": "hs",
    "java": "java",
    "javascript": "js",
    "lua": "lua",
    "perl": "pl",
    "php": "php",
    "python": "py",
    "ruby": "rb",
    "rust": "rs",
    "scala": "scala",
    "plaintext": "txt",
}

var fileExtAdditional = map[string][]string{
    "c": []string{"h"},
}

func DefaultFname(lang string) (string, bool) {
    ext, ok := fileExt[lang]

    if lang == "java" {
        return "Main." + ext, ok
    }

    return "main." + ext, ok
}

func DefaultContent(lang string) (string, bool) {
    content, ok := defaultContent[lang]
    return content, ok
}

func DetermineLanguage(path string) (string, bool) {

    if len(filepath.Ext(path)) != 0 {
        extension := filepath.Ext(path)[1:]

        for lang, ext := range fileExt {
            if ext == extension {
                return lang, true
            }
        }
    }

    return "", false
}

func AllowedExtensions(lang string) []string {
    var extensions = make([]string, 0)

    ext, ok := fileExt[lang]
    if ok {
        extensions = append(extensions, ext)
    }

    extAdditional, ok := fileExtAdditional[lang]
    if ok {
        extensions = append(extensions, extAdditional...)
    }

    return extensions
}

func AllowedExtension(ext, lang string) bool {
    allowed := AllowedExtensions(lang)
    for _, extension := range allowed {
        if extension == ext {
            return true
        }
    }

    return false
}
