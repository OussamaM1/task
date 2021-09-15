package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do command called")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
