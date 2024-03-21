// Package cmd provides functionality related to the command line.
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as complete, in-progress",
	Run:   do,
}

func init() {
	RootCommand.AddCommand(doCommand)
}

func do(_ *cobra.Command, args []string) {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Cannot convert the argument : ", arg)
		} else {
			ids = append(ids, id)
		}
	}
	fmt.Println(ids)
}
