package main

import "github.com/oussamaM1/task/cmd"

func main() {
	if err := cmd.RootCommand.Execute(); err != nil {
		return
	}
}
