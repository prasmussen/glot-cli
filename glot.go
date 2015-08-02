package main

import (
    "os"
    "fmt"
    "path/filepath"
    "./util"
    "./config"
    "./cli"
    "./snippets"
    "./run"
)

const (
    Version = "1.1.0"
)

var (
	AppPath = filepath.Join(util.Homedir(), ".glot")
)

func main() {
    cli.AddHandler("new <language>", newSnippet, "Create new snippet")
    cli.AddHandler("run <path>", runLatest, "Run code")
    cli.AddHandler("run <path> --version <version>", runVersion, "Run code code with a specific language version")
    cli.AddHandler("publish --title <title>", publishSnippet, "Publish snippet")
    cli.AddHandler("languages", listLanguages, "List available languages available to run")
    cli.AddHandler("versions <language>", listVersions, "List available versions for a language")
    cli.AddHandler("help", printHelp, "Print help")
    cli.AddHandler("--version", printAppVersion, "Print application version")

    ok := cli.Handle(os.Args[1:])
    if !ok {
        fmt.Fprintf(os.Stderr, "No valid arguments given, use '%s help' to see available commands\n", util.AppName())
        os.Exit(1)
    }
}

func newSnippet(args map[string]string) {
    snippets.NewSnippet(args["language"])
}

func publishSnippet(args map[string]string) {
    cfg := getConfig()
    snippets.Publish(cfg, args["title"])
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

func printAppVersion(args map[string]string) {
    fmt.Printf("%s v%s\n", util.AppName(), Version)
}

func printHelp(args map[string]string) {
    appName := util.AppName()
    fmt.Printf("%s usage:\n\n", appName)

    for _, h := range cli.GetHandlers() {
        fmt.Printf("%s %s  (%s)\n", appName, h.Pattern, h.Description)
    }
}

func getConfig() *config.Config {
    configPath := filepath.Join(AppPath, "config")
    return readOrCreateConfig(configPath)
}

func readOrCreateConfig(path string) *config.Config {
    if util.FileExists(path) {
        // Open existing config
        cfg, err := config.ReadConfig(path)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to read config: %s\n", err.Error())
            os.Exit(1)
        }

        return cfg
    } else {
        // Create new config
        fmt.Println("Config not found, creating new...")
        token := util.PromptInput("Enter API token: ")
        cfg := config.DefaultConfig()
        cfg.Run.Token = token
        cfg.Snippets.Token = token

        // Save config
        err := config.SaveConfig(path, cfg)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to save config: %s\n", err.Error())
            os.Exit(1)
        }
        fmt.Printf("Config saved to %s\n\n", path)

        return cfg
    }
}
