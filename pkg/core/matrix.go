package core

import "fmt"

// Matrix represents a simple 2D matrix structure
type Matrix struct {
	Rows, Cols int
	Data       [][]float64
}

// NewMatrix initializes a matrix
func NewMatrix(rows, cols int) *Matrix {
	fmt.Println("Creating new matrix")
	return &Matrix{Rows: rows, Cols: cols}
}
