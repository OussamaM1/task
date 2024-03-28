// Package services provides functionality related to configuration and managing todos file.
package services

import (
	"log"
	"os"
)

// fileName const - is the name of the todos file.
const fileName string = "todos.txt"

// createFile func - Created the todos file if not found
func createFile() *os.File {
	LogInfo("Creating a new file. \n")
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
	LogInfo("Open Todos file.")
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return createFile()
	}
	return file
}

// ReadFile func - Reads the content of todos file
func ReadFile() []byte {
	LogInfo("Reading Todos file.")
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Could not read data from todos file, ", err)
	}
	return data
}

// WriteFile func - Write the todos file
func WriteFile(data []byte) {
	LogInfo("Writing data into todos file.")
	// Open the file
	file := OpenFile()
	// Close the file when the function returns
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalf("Error while closing the file: %v", err)
		}
	}(file)
	// Write data to the file
	length, err := file.WriteString(string(data))
	if err != nil {
		log.Fatalf("Failed writing to file: %v", err)
	}
	LogInfo("File Name: %s", file.Name())
	LogInfo("Length: %d bytes", length)
}

// OverwriteFile func - Empty file content
func OverwriteFile() {
	file := OpenFile()
	err := file.Truncate(0)
	if err != nil {
		log.Fatalln("Error truncating file:", err)
	}
}
