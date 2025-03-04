package tests

import (
	"testing"
	"goai/models"
)

func TestCustomModelTraining(t *testing.T) {
	model := models.CustomModel{Name: "TestModel"}
	model.Train() // Should not return an error

	// Placeholder test since training doesn't return a value
	if model.Name == "" {
		t.Errorf("Expected model name to be set")
	}
}
