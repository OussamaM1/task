// Package cmd provides functionality related to the command line.
package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// RootCommand Command - is the root command for the application.
var RootCommand = &cobra.Command{
	Use:   "task",
	Short: "📌 task is golang-based Command-line interface designed for efficient management of your TODOs.",
}

// Execute func - runs the root command.
func Execute() {
	if err := RootCommand.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
