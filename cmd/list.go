package cmd

import (
	"fmt"
	"github.com/oussamaM1/task/models"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"log"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run:   list,
}

func init() {
	RootCommand.AddCommand(listCommand)
}

func list(cmd *cobra.Command, args []string) {
	var todoList []models.Todo
	todoList = append(todoList, models.Todo{Task: "This is a test task", State: "In-progress"})
	todoList = append(todoList, models.Todo{Task: "Second task", State: "In-progress"})
	yamlData, err := yaml.Marshal(todoList)
	if err != nil {
		log.Fatalf("Error marshalling TODOs to YAML: %v", err)
	}
	fmt.Printf("List of TODOs:\n%s\n", yamlData)
}
