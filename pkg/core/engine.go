package core

import (
    "fmt"
    "log"
)

type Engine struct {
    Model  *Model
    Config Config
}

func NewEngine(cfg Config) *Engine {
    return &Engine{
        Model:  &Model{Name: "default", Weights: []float64{0.1, 0.2}, Bias: 0.0},
        Config: cfg,
    }
}

func (e *Engine) Train(data []DataPoint) error {
    if len(data) == 0 {
        return fmt.Errorf("no training data provided")
    }

    for epoch := 0; epoch < e.Config.Epochs; epoch++ {
        totalError := 0.0
        for _, dp := range data {
            pred, err := e.Predict(dp.Features)
            if err != nil {
                return err
            }
            error := dp.Label - pred

            for i := range e.Model.Weights {
                e.Model.Weights[i] += e.Config.LearningRate * error * dp.Features[i]
            }
            e.Model.Bias += e.Config.LearningRate * error

            totalError += error * error
        }
        log.Printf("Epoch %d: Error = %.6f", epoch, totalError/float64(len(data)))
    }

    if e.Config.Blockchain {
        if err := e.storeModelOnChain(); err != nil {
            return fmt.Errorf("blockchain storage failed: %v", err)
        }
    }
    return nil
}
