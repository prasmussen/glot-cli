package main

import (
    "os"
    "./config"
    "./cli"
    "./snippets"
    "./run"
)

func main() {
    cli.AddHandler("new <language>", newSnippet, "Create new snippet")
    cli.AddHandler("snippets", listSnippets, "List snippets")
    //cli.AddHandler("snippets --page <n>", listSnippets, "List snippets")
    cli.AddHandler("meta <id>", printMetaSnippet, "Print snippet meta information")
    cli.AddHandler("clone <id>", cloneSnippet, "Clone snippet into a directory")
    cli.AddHandler("pull", pullSnippet, "Pull snippet from api, will overwrite local changes")
    cli.AddHandler("push", pushSnippet, "Push snippet to api, will overwrite remote changes")
    cli.AddHandler("languages", listLanguages, "List available languages available to run")
    cli.AddHandler("versions <language>", listVersions, "List available versions for a language")
    cli.AddHandler("run <path>", runLatest, "Run code")
    cli.AddHandler("run <path> --version <version>", runVersion, "Run code code with a specific language version")
    cli.Handle(os.Args[1:])
}

func getConfig() *config.Config {
    return config.ReadConfig("config.json")
}

func newSnippet(args map[string]string) {
    snippets.NewSnippet(args["language"])
}

func printMetaSnippet(args map[string]string) {
    cfg := getConfig()
    snippets.PrintMeta(cfg, args["id"])
}

func cloneSnippet(args map[string]string) {
    cfg := getConfig()
    snippets.Clone(cfg, args["id"])
}

func pullSnippet(args map[string]string) {
    cfg := getConfig()
    snippets.Pull(cfg)
}

func pushSnippet(args map[string]string) {
    cfg := getConfig()
    snippets.Push(cfg)
}

func listSnippets(args map[string]string) {
    cfg := getConfig()
    snippets.ListSnippets(cfg)
}

func listLanguages(args map[string]string) {
    cfg := getConfig()
    run.ListLanguages(cfg)
}

func listVersions(args map[string]string) {
    cfg := getConfig()
    run.ListVersions(cfg, args["language"])
}

func runLatest(args map[string]string) {
    cfg := getConfig()
    run.Run(cfg, "latest", args["path"])
}

func runVersion(args map[string]string) {
    cfg := getConfig()
    run.Run(cfg, args["version"], args["path"])
}
