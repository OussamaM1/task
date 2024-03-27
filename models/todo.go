// Package models provides functionality and types related to models.
package models

// Task struct - Contains ID, Description and State
type Task struct {
	ID          int    `yaml:"Id"`
	Description string `yaml:"Description"`
	State       string `yaml:"State"`
}
