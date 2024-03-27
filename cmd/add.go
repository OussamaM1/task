// Package cmd provides functionality related to the command line.
package cmd

import (
	"github.com/oussamaM1/task/models"
	"github.com/oussamaM1/task/services"
	"github.com/spf13/cobra"
	"strings"
)

var addCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run:   add,
}

func init() {
	RootCommand.AddCommand(addCommand)
}

func add(cmd *cobra.Command, args []string) {
	services.LogInfo("Command called: %s", cmd.Name())
	services.LogInfo("Arguments: %s", strings.Join(args, " "))
	var id int = services.GenerateNewID()
	var taskList []models.Task
	for _, arg := range args {
		model := models.Task{ID: id, Description: arg, State: "In-progress"}
		taskList = append(taskList, model)
		id++
	}
	services.WriteData(taskList)
}
