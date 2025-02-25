package main

import (
    "fmt"
    "os"
    "github.com/yourorg/goai/pkg/core"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: goai <command> [args]")
        fmt.Println("Commands: hello, train, predict")
        os.Exit(1)
    }

    eng := core.NewEngine(core.Config{
        LearningRate: 0.01,
        Epochs:       10,
        Blockchain:   true,
    })

    switch os.Args[1] {
    case "hello":
        fmt.Println(core.HelloAI())
    case "train":
        data := []core.DataPoint{
            {Features: []float64{1.0, 2.0}, Label: 1.0},
            {Features: []float64{2.0, 1.0}, Label: 0.0},
        }
        if err := eng.Train(data); err != nil {
            fmt.Println("Training error:", err)
        } else {
            fmt.Println("Training complete!")
        }
    case "predict":
        features := []float64{1.5, 1.5}
        pred, err := eng.Predict(features)
        if err != nil {
            fmt.Println("Prediction error:", err)
        } else {
            fmt.Printf("Prediction: %.4f\n", pred)
        }
    default:
        fmt.Println("Unknown command:", os.Args[1])
    }
}
