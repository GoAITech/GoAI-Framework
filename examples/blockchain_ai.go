package main

import (
	"fmt"
	"goai/pkg/blockchain"
)

func main() {
	fmt.Println("Initializing AI-powered blockchain module...")

	// Example: Deploy AI model on a blockchain smart contract
	chain := blockchain.NewBlockchain()
	modelID := chain.DeployModel("example_model")

	fmt.Printf("Model deployed with ID: %s\n", modelID)
}
