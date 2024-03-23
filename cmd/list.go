// Package cmd provides functionality related to the command line.
package cmd

import (
	"github.com/oussamaM1/task/services"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run:   list,
}

func init() {
	RootCommand.AddCommand(listCommand)
}

func list(_ *cobra.Command, _ []string) {
	services.Logf("List command called")
}
