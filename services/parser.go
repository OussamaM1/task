package services

import (
	"github.com/oussamaM1/task/models"
	"gopkg.in/yaml.v2"
	"log"
	"strconv"
)

// WriteData func - responsible for parsing data in yaml format
func WriteData(taskList []models.Task) {
	yamlData, err := yaml.Marshal(taskList)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}
	WriteFile(yamlData)
}

// ReadData func - responsible for reading data in yaml format
func ReadData() []models.Task {
	var data = ReadFile()
	var taskList []models.Task
	err := yaml.Unmarshal(data, &taskList)
	if err != nil {
		log.Fatalf("Failed to unmarshal data: %v", err)
	}
	return taskList
}

// GenerateNewID func - responsible for generating a unique ID for each new task added
func GenerateNewID() int {
	taskList := ReadData()
	if taskList == nil {
		return 1
	}
	return taskList[len(taskList)-1].ID + 1
}

// EditTask func - edits a task and parse the list of tasks updated
func EditTask(args []string) {
	if len(args) != 2 {
		log.Fatalln("The correct format is task do <ID of the task> <State -Complete, In-Progress, Open->")
	}
	taskList := ReadData()
	index, err := strconv.Atoi(args[0])
	if err != nil || (index < 1 || index > len(taskList)) {
		log.Fatalln("Provide a valid ID task.")
	}
	if args[1] != "Open" && args[1] != "In-progress" && args[1] != "Completed" {
		log.Fatalln("Invalid state detected. Expected state values are 'Open', 'In-progress', or 'Completed', but the actual state provided is: ", args[1])
	}
	taskList[index-1].State = args[1]
	OverwriteFile()
	WriteData(taskList)
}
