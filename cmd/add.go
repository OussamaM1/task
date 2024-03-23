// Package cmd provides functionality related to the command line.
package cmd

import (
	"github.com/oussamaM1/task/services"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run:   add,
}

func init() {
	RootCommand.AddCommand(addCommand)
}

func add(_ *cobra.Command, _ []string) {
	services.Logf("add command called")
}
