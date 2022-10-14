//go:build windows

package main

import "os"

func GetDefaultUserHomePath() string {
	dir, _ := os.UserHomeDir()
	return dir
}
