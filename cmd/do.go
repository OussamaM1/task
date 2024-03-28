// Package cmd provides functionality related to the command line.
package cmd

import (
	"fmt"
	"github.com/oussamaM1/task/services"
	"github.com/spf13/cobra"
	"strings"
)

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete, in-progress and Open",
	Run:   do,
}

func init() {
	RootCommand.AddCommand(doCommand)
}

func do(cmd *cobra.Command, args []string) {
	services.LogInfo("Command called: %s", cmd.Name())
	services.LogInfo("Arguments: %s", strings.Join(args, " "))
	services.EditTask(args)
	fmt.Println("✅ The task has been successfully updated.")
}
