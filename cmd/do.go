// Package cmd provides functionality related to the command line.
package cmd

import (
	"github.com/oussamaM1/task/services"
	"github.com/spf13/cobra"
)

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete, in-progress",
	Run:   do,
}

func init() {
	RootCommand.AddCommand(doCommand)
}

func do(_ *cobra.Command, _ []string) {
	services.LogInfo("do command called")
}
