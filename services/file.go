package services

import (
	"fmt"
	"log"
	"os"
)

func ReadFile() []byte {
	data, err := os.ReadFile("todos")
	if err != nil {
		createFile()
	}
	return data
}

func createFile() {

	// todo: Replace with logs
	fmt.Printf("Creating a new file: \n")

	file, err := os.Create("todos")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	defer func(file *os.File) {
		g
		err := file.Close()
		if err != nil {
			log.Fatalf("Error while closing the file: %s", err)
		}
	}(file)

	length, err := file.WriteString("This is a test new file")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", length)
}
func WriteFile() {

}

func OverwriteFile() {

}
