package services

import (
	"github.com/oussamaM1/task/models"
	"gopkg.in/yaml.v2"
	"log"
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
