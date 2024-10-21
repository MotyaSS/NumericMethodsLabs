package main

import (
	"Chislaki/gauss_algo"
	"Chislaki/thomas_algo"
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func PrintBigRatSlice(slice []big.Rat) {
	for _, rat := range slice {
		fmt.Print(rat.String(), " ")
	}
	fmt.Println()
}

// ReadMatrixFromFile reads a matrix and corresponding values from a file.
func ReadMatrixFromFile(filename string) ([][]big.Rat, []big.Rat, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var matrix [][]big.Rat
	var values []big.Rat

	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Split(line, ",")
		var row []big.Rat
		for _, numStr := range numStrs {
			numStr = strings.TrimSpace(numStr)
			num := new(big.Rat)
			if _, ok := num.SetString(numStr); !ok {
				return nil, nil, fmt.Errorf("invalid number format: %s", numStr)
			}
			row = append(row, *num)
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	// Separate the last column as the values
	for i := range matrix {
		values = append(values, matrix[i][len(matrix[i])-1])
		matrix[i] = matrix[i][:len(matrix[i])-1]
	}

	return matrix, values, nil
}

// Thomas Algorithm example
func exampleThomas() {
	solution, determinant, err := thomas_algo.ThomasFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("solution: ")
	PrintBigRatSlice(solution)
	fmt.Println("determinant: ")
	fmt.Println(determinant.String())
}

// IsInverse checks if matrix B is the inverse of matrix A
func IsInverse(A, B [][]big.Rat) (bool, error) {
	dim := len(A)
	if dim == 0 || len(B) != dim || len(A[0]) != dim || len(B[0]) != dim {
		return false, fmt.Errorf("matrices are not square or dimensions do not match")
	}

	// Create an identity matrix
	identity := make([][]big.Rat, dim)
	for i := range identity {
		identity[i] = make([]big.Rat, dim)
		for j := range identity[i] {
			if i == j {
				identity[i][j] = *big.NewRat(1, 1)
			} else {
				identity[i][j] = *big.NewRat(0, 1)
			}
		}
	}

	// Multiply A and B
	product := make([][]big.Rat, dim)
	for i := range product {
		product[i] = make([]big.Rat, dim)
		for j := range product[i] {
			product[i][j] = *big.NewRat(0, 1)
			for k := 0; k < dim; k++ {
				temp := new(big.Rat).Mul(&A[i][k], &B[k][j])
				product[i][j].Add(&product[i][j], temp)
			}
		}
	}

	// Check if the product is the identity matrix
	for i := range product {
		for j := range product[i] {
			if product[i][j].Cmp(&identity[i][j]) != 0 {
				return false, nil
			}
		}
	}
	return true, nil
}

// Gaussian Elimination example
func exampleGauss() {
	matrix, values, err := ReadMatrixFromFile("test2.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Matrix and values printed out
	fmt.Println("Matrix:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Print(val.String(), " ")
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Values:")
	for _, val := range values {
		fmt.Print(val.String(), " ")
	}
	fmt.Println()

	// Some actual solving
	solution, _, inverse, err := gauss_algo.GaussianElimination(matrix, values)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	isInverse, err := IsInverse(matrix, inverse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if isInverse {
		fmt.Println("Inverse is correct")
		fmt.Println("Inverse:")
		for _, row := range inverse {
			PrintBigRatSlice(row)
		}
	} else {
		fmt.Println("Inverse is incorrect")
	}
	fmt.Println()

	fmt.Println("Solution:")
	PrintBigRatSlice(solution)
	fmt.Println()
}

func main() {
	exampleGauss()
}
