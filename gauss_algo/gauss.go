package gauss_algo

import (
	"fmt"
	"math/big"
)

func printBigRatSlice(slice []big.Rat) {
	for _, rat := range slice {
		fmt.Print(rat.String(), " ")
	}
	fmt.Println()
}

// InvertMatrix performs Gaussian-Jordan elimination to find the inverse of a matrix.
func InvertMatrix(matrix [][]big.Rat) ([][]big.Rat, error) {
	dim := len(matrix)
	if dim == 0 {
		return nil, fmt.Errorf("matrix is empty")
	}

	// Create an augmented matrix with the identity matrix
	augmentedMatrix := make([][]big.Rat, dim)
	for i := range matrix {
		augmentedMatrix[i] = make([]big.Rat, 2*dim)
		for j := range matrix[i] {
			augmentedMatrix[i][j] = *new(big.Rat).Set(&matrix[i][j])
		}
		for j := dim; j < 2*dim; j++ {
			if j-dim == i {
				augmentedMatrix[i][j] = *big.NewRat(1, 1)
			} else {
				augmentedMatrix[i][j] = *big.NewRat(0, 1)
			}
		}
	}

	// Perform Gaussian-Jordan elimination
	for i := 0; i < dim; i++ {
		// Find the pivot
		pivot := augmentedMatrix[i][i]
		if pivot.Cmp(big.NewRat(0, 1)) == 0 {
			return nil, fmt.Errorf("matrix is singular")
		}

		// Normalize the pivot row
		for j := 0; j < 2*dim; j++ {
			augmentedMatrix[i][j].Quo(&augmentedMatrix[i][j], &pivot)
		}

		// Eliminate the current column in the rows above and below
		for k := 0; k < dim; k++ {
			if k != i {
				factor := new(big.Rat).Set(&augmentedMatrix[k][i])
				for j := 0; j < 2*dim; j++ {
					var temp *big.Rat
					temp = new(big.Rat).Mul(factor, &augmentedMatrix[i][j])
					augmentedMatrix[k][j].Sub(&augmentedMatrix[k][j], temp)
				}
			}
		}

	}

	// Extract the inverse matrix
	inverse := make([][]big.Rat, dim)
	for i := range inverse {
		inverse[i] = make([]big.Rat, dim)
		for j := 0; j < dim; j++ {
			inverse[i][j] = augmentedMatrix[i][j+dim]
		}
	}

	return inverse, nil
}

// GaussianElimination performs Gaussian elimination and returns the solution, determinant, and inverse matrix.
func GaussianElimination(matrix [][]big.Rat, values []big.Rat) ([]big.Rat, big.Rat, [][]big.Rat, error) {
	dim := len(matrix)
	if dim == 0 {
		return nil, big.Rat{}, nil, fmt.Errorf("matrix is empty")
	}

	// Augment the matrix with the values
	augmentedMatrix := make([][]big.Rat, dim)
	for i := range matrix {
		augmentedMatrix[i] = make([]big.Rat, len(matrix[i])+1)
		for j, val := range matrix[i] {
			augmentedMatrix[i][j] = *new(big.Rat).Set(&val)
		}
		augmentedMatrix[i][len(matrix[i])] = values[i]
	}

	// Perform Gaussian elimination to find the solution
	for i := 0; i < dim; i++ {
		// Find the pivot
		pivot := augmentedMatrix[i][i]
		if pivot.Cmp(big.NewRat(0, 1)) == 0 {
			return nil, big.Rat{}, nil, fmt.Errorf("matrix is singular")
		}

		// Normalize the pivot row
		for j := 0; j <= dim; j++ {
			augmentedMatrix[i][j].Quo(&augmentedMatrix[i][j], &pivot)
		}

		// Eliminate the current column in the rows above and below
		for k := 0; k < dim; k++ {
			if k != i {
				factor := augmentedMatrix[k][i]
				for j := 0; j <= dim; j++ {
					temp := new(big.Rat).Mul(&factor, &augmentedMatrix[i][j])
					augmentedMatrix[k][j].Sub(&augmentedMatrix[k][j], temp)
				}
			}
		}
	}

	// Extract the solution vector
	solution := make([]big.Rat, dim)
	for i := 0; i < dim; i++ {
		solution[i] = augmentedMatrix[i][dim]
	}

	// Calculate the determinant
	determinant := *big.NewRat(1, 1)
	for i := 0; i < dim; i++ {
		determinant.Mul(&determinant, &augmentedMatrix[i][i])
	}

	// Inverse Matrix calculation
	inverse, err := InvertMatrix(matrix)
	if err != nil {
		return nil, big.Rat{}, nil, err
	}

	return solution, determinant, inverse, nil
}
