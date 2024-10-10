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

func PrintBigRatSlice(slice []*big.Rat) {
	for _, rat := range slice {
		fmt.Print(rat.String(), " ")
	}
	fmt.Println()
}

func PrintBigRatSlice2(slice []big.Rat) {
	for _, rat := range slice {
		fmt.Print(rat.String(), " ")
	}
	fmt.Println()
}

// ReadMatrixFromFile reads a matrix and corresponding values from a file.
func ReadMatrixFromFile(filename string) ([][]*big.Rat, []*big.Rat, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]*big.Rat
	var values []*big.Rat

	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Split(line, ",")
		var row []*big.Rat
		for _, numStr := range numStrs {
			numStr = strings.TrimSpace(numStr)
			num := new(big.Rat)
			if _, ok := num.SetString(numStr); !ok {
				return nil, nil, fmt.Errorf("invalid number format: %s", numStr)
			}
			row = append(row, num)
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

func exampleThomas() {
	solution, determinant, err := thomas_algo.ThomasFile("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("solution: ")
	PrintBigRatSlice2(solution)
	fmt.Println("determinant: ")
	fmt.Println(determinant.String())
}

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
	fmt.Println("Inverse:")
	for _, row := range inverse {
		PrintBigRatSlice(row)
	}
	fmt.Println("Solution:")
	PrintBigRatSlice(solution)
}

func main() {
	exampleGauss()
}
