package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var RootCommand = &cobra.Command{
	Use:   "task",
	Short: "📌 task is golang-based Command-line interface designed for efficient management of your TODOs.",
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
