package main

import (
	"fmt"
	"goai/core/engine"
)

func main() {
	fmt.Println("Welcome to GoAI!")

	// Example of running a simple AI function
	result := engine.SimpleComputation(5, 10)
	fmt.Printf("Computation result: %d\n", result)
}
