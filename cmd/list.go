package cmd

import (
	"fmt"
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

func list(cmd *cobra.Command, args []string) {
	fmt.Println("list command called")
}
