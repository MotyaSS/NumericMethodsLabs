package equation_system_solve_algo

import (
	"math"
)

// Matrix represents a mathematical matrix
type Matrix struct {
	data [][]float64
	rows int
	cols int
}

// NewMatrix creates a new matrix with given dimensions
func NewMatrix(rows, cols int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return &Matrix{
		data: data,
		rows: rows,
		cols: cols,
	}
}

// Get returns the value at given position
func (m *Matrix) Get(i, j int) float64 {
	return m.data[i][j]
}

// Set sets the value at given position
func (m *Matrix) Set(i, j int, value float64) {
	m.data[i][j] = value
}

// Subtract subtracts another matrix from this matrix
func (m *Matrix) Subtract(other *Matrix) *Matrix {
	if m.rows != other.rows || m.cols != other.cols {
		panic("Matrix dimensions do not match for subtraction")
	}
	result := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[i][j] = m.data[i][j] - other.data[i][j]
		}
	}
	return result
}

// Multiply multiplies this matrix by another matrix
func (m *Matrix) Multiply(other *Matrix) *Matrix {
	if m.cols != other.rows {
		panic("Matrix dimensions do not match for multiplication")
	}
	result := NewMatrix(m.rows, other.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			sum := 0.0
			for k := 0; k < m.cols; k++ {
				sum += m.data[i][k] * other.data[k][j]
			}
			result.data[i][j] = sum
		}
	}
	return result
}

// VecNormC calculates the vector norm (maximum absolute value)
func (m *Matrix) VecNormC() float64 {
	if m.cols != 1 {
		panic("VecNormC can only be calculated for column vectors")
	}
	max := 0.0
	for i := 0; i < m.rows; i++ {
		abs := math.Abs(m.data[i][0])
		if abs > max {
			max = abs
		}
	}
	return max
}

// MatrixNormC calculates the matrix norm (maximum absolute row sum)
func (m *Matrix) MatrixNormC() float64 {
	max := 0.0
	for i := 0; i < m.rows; i++ {
		sum := 0.0
		for j := 0; j < m.cols; j++ {
			sum += math.Abs(m.data[i][j])
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// InverseMatrix calculates the inverse of a 2x2 matrix
func (m *Matrix) InverseMatrix() *Matrix {
	if m.rows != 2 || m.cols != 2 {
		panic("InverseMatrix is only implemented for 2x2 matrices")
	}
	det := m.data[0][0]*m.data[1][1] - m.data[0][1]*m.data[1][0]
	if math.Abs(det) < 1e-10 {
		panic("Matrix is singular")
	}
	result := NewMatrix(2, 2)
	result.data[0][0] = m.data[1][1] / det
	result.data[0][1] = -m.data[0][1] / det
	result.data[1][0] = -m.data[1][0] / det
	result.data[1][1] = m.data[0][0] / det
	return result
}
