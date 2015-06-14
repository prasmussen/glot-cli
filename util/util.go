package util

import (
    "runtime"
    "os"
)

// Returns the users home dir
func Homedir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("APPDATA")
	}
	return os.Getenv("HOME")
}
