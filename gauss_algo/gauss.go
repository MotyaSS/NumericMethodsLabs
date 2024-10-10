package gauss_algo

import (
	"fmt"
	"math/big"
)

// GaussianElimination performs Gaussian elimination and returns the solution, determinant, and inverse matrix.
func GaussianElimination(matrix [][]*big.Rat, values []*big.Rat) ([]*big.Rat, *big.Rat, [][]*big.Rat, error) {
	dim := len(matrix)
	if dim == 0 {
		return nil, nil, nil, fmt.Errorf("matrix is empty")
	}

	// Augment the matrix with the identity matrix
	augmentedMatrix := make([][]*big.Rat, dim)
	for i := range matrix {
		augmentedMatrix[i] = append(matrix[i], make([]*big.Rat, dim)...)
		for j := 0; j < dim; j++ {
			if i == j {
				augmentedMatrix[i][dim+j] = big.NewRat(1, 1)
			} else {
				augmentedMatrix[i][dim+j] = big.NewRat(0, 1)
			}
		}
	}

	// Perform Gaussian elimination
	for i := 0; i < dim; i++ {
		// Find the pivot
		pivot := augmentedMatrix[i][i]
		if pivot.Cmp(big.NewRat(0, 1)) == 0 {
			return nil, nil, nil, fmt.Errorf("matrix is singular")
		}

		// Normalize the pivot row
		for j := 0; j < 2*dim; j++ {
			augmentedMatrix[i][j].Quo(augmentedMatrix[i][j], pivot)
		}

		// Eliminate the current column in the rows above and below
		for k := 0; k < dim; k++ {
			if k != i {
				factor := augmentedMatrix[k][i]
				for j := 0; j < 2*dim; j++ {
					augmentedMatrix[k][j].Sub(augmentedMatrix[k][j], new(big.Rat).Mul(factor, augmentedMatrix[i][j]))
				}
			}
		}
	}

	// Extract the inverse matrix
	inverse := make([][]*big.Rat, dim)
	for i := range inverse {
		inverse[i] = make([]*big.Rat, dim)
		for j := 0; j < dim; j++ {
			inverse[i][j] = augmentedMatrix[i][dim+j]
		}
	}

	// Extract the solution vector
	solution := make([]*big.Rat, dim)
	for i := 0; i < dim; i++ {
		solution[i] = augmentedMatrix[i][dim*2-1]
	}

	// Calculate the determinant
	determinant := big.NewRat(1, 1)
	for i := 0; i < dim; i++ {
		determinant.Mul(determinant, augmentedMatrix[i][i])
	}

	return solution, determinant, inverse, nil
}
