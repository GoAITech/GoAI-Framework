package main

import (
	"fmt"
	"github.com/yourorg/goai"
)

func main() {
	model := goai.NewModel()
	result := model.Predict([]float64{1.2, 3.4, 5.6})
	fmt.Println("AI Prediction:", result)
}
