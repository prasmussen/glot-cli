package util

import (
    "runtime"
    "os"
    "fmt"
)

// Returns the users home dir
func Homedir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("APPDATA")
	}
	return os.Getenv("HOME")
}

// Prompt user to input data
func PromptInput(msg string) string {
    fmt.Printf(msg)
    var str string
    fmt.Scanln(&str)
    return str
}
