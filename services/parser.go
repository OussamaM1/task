package services

import (
	"github.com/oussamaM1/task/models"
	"gopkg.in/yaml.v2"
	"log"
)

// WriteData func - responsible for parsing data in yaml format
func WriteData(todoList []models.Todo) {
	yamlData, err := yaml.Marshal(todoList)
	if err != nil {
		log.Fatalf("Failed to marshal YAML: %v", err)
	}
	WriteFile(yamlData)
}
