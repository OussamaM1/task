// Package models provides functionality and types related to models.
package models

// Todo struct - Contains two variables Task and State
type Todo struct {
	Task  string `yaml:"task"`
	State string `yaml:"state"`
}
