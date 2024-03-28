// Package services provides functionality related to configuration and managing todos file.
package services

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/oussamaM1/task/models"
	"os"
)

// PrintTasks func - Prints the current tasks in form of table
func PrintTasks(taskList []models.Task) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Description", "State"})
	t.AppendSeparator()
	for _, task := range taskList {
		t.AppendRow([]interface{}{task.ID, task.Description, stateColor(task.State)})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}

// stateColor func - Give a custom color output to the state of task
func stateColor(state string) string {
	switch state {
	case "In-progress":
		return fmt.Sprintf("%s%s%s\n", colorBlue, state, colorReset)
	case "Completed":
		return fmt.Sprintf("%s%s%s\n", colorGreen, state, colorReset)
	case "Open":
		return fmt.Sprintf("%s%s%s\n", colorGray, state, colorReset)
	}
	return state
}
