// Package main - the entry point for the application.
package main

import (
	"github.com/oussamaM1/task/cmd"
	"github.com/oussamaM1/task/services"
)

func main() {
	services.Logf("Running application in development mode.\n")
	services.ReadFile()
	cmd.Execute()
}
