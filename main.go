package main

import (
	"github.com/oussamaM1/task/cmd"
	"github.com/oussamaM1/task/services"
)

func main() {
	// todo: Open new File if todos not found
	services.ReadFile()
	cmd.Execute()
}
