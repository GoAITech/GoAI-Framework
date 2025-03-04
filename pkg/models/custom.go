package models

import "fmt"

// CustomModel represents user-defined AI models
type CustomModel struct {
	Name string
}

// Train runs training for the custom model
func (m *CustomModel) Train() {
	fmt.Printf("Training custom model: %s\n", m.Name)
}
