// Package services provides functionality related to configuration and managing todos file.
package services

import (
	"github.com/oussamaM1/task/models"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

// fileName const - is the name of the todos file.
const fileName string = "todos.txt"

// createFile func - Created the todos file if not found
func createFile() *os.File {
	Logf("Creating a new file. \n")
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Failed creating file: ", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln("Error while closing the file: ", err)
		}
	}(file)
	return file
}

// OpenFile func - Opens the todos file
func OpenFile() *os.File {
	Logf("Open Todos file.")
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return createFile()
	}
	return file
}

// ReadFile func - Reads the content of todos file
func ReadFile() []byte {
	Logf("Reading Todos file.")
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Could not read data from todos file, ", err)
	}
	return data
}

// WriteFile func - Write the todos file
func WriteFile() {
	Logf("Writing data into todos file.")
	// Open the file
	file := OpenFile()
	// Close the file when the function returns
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalf("Error while closing the file: %v", err)
		}
	}(file)
	var todoList []models.Todo
	todoList = append(todoList, models.Todo{Task: "This is a test task", State: "In-progress"})
	todoList = append(todoList, models.Todo{Task: "Second task", State: "In-progress"})
	yamlData, err := yaml.Marshal(todoList)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}
	// Write data to the file
	length, err := file.WriteString(string(yamlData))
	if err != nil {
		log.Fatalf("Failed writing to file: %v", err)
	}
	Logf("File Name: %s", file.Name())
	Logf("Length: %d bytes", length)
	Logf("File Content: \n%s ", string(ReadFile()))
}
