// Package cmd provides functionality related to the command line.
package cmd

import (
	"fmt"
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

func list(cmd *cobra.Command, _ []string) {
	services.LogInfo("Command called: %s", cmd.Name())
	services.LogInfo("List of tasks: ")
	taskList := services.ReadData()
	if taskList == nil {
		fmt.Println("⚠️ There are currently no tasks available. You can use the \"add\" command to include a new task.")
	} else {
		services.PrintTasks(taskList)
	}
}
