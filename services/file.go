package services

import (
	"log"
	"os"
)

// createFile func - Created the todos file if not found
func createFile() *os.File {
	Logf("Creating a new file. \n")
	file, err := os.Create("todos")
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

// ReadFile func - Reads the todos file
func ReadFile() *os.File {
	Logf("Reading Todos file. \n")
	_, err := os.ReadFile("todos")
	if err != nil {
		return createFile()
	}
	return nil
}

// WriteFile func - Write the todos file
func WriteFile() {
	file := ReadFile()
	length, err := file.WriteString("This is a test new file /n")
	if err != nil {
		log.Fatalln("Failed writing to file: ", err)
	}
	Logf("\nFile Name: %s", file.Name())
	Logf("\nLength: %d bytes", length)
}

func OverwriteFile() {

}
