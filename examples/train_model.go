package main

import (
	"fmt"
	"goai/core/models"
)

func main() {
	fmt.Println("Training AI Model...")

	// Load dataset
	data := models.LoadDataset("dataset.csv")

	// Train model
	model := models.Train(data)

	// Save model
	model.Save("trained_model.ai")

	fmt.Println("Model training complete!")
}
