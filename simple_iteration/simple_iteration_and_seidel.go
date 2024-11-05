package simple_iteration_algo

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadMatrix(filename string) ([][]float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	var matrix [][]float64
	var values []float64
	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Split(line, ",")
		var row []float64
		for _, numStr := range numStrs {
			numStr = strings.TrimSpace(numStr)
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
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

// region Iterations

func calcCurIterX(alpha [][]float64, beta []float64, prevIterX []float64) []float64 {
	dim := len(alpha)
	curIterX := make([]float64, dim)
	for i := range curIterX {
		curIterX[i] = beta[i]
		for j := range alpha[i] {
			curIterX[i] += alpha[i][j] * prevIterX[j]
		}
	}
	return curIterX
}

func calcSeidelCurIterX(alpha [][]float64, beta []float64, prevIterX []float64) []float64 {
	dim := len(alpha)
	curIterX := make([]float64, dim)
	for i := range curIterX {
		curIterX[i] = beta[i]
		for j := 0; j < i; j++ {
			curIterX[i] += alpha[i][j] * curIterX[j]
		}
		for j := i; j < dim; j++ {
			curIterX[i] += alpha[i][j] * prevIterX[j]
		}

	}
	return curIterX
}

func calcAlphaBeta(matrix [][]float64, values []float64) ([][]float64, []float64) {
	dim := len(matrix)
	beta := make([]float64, dim)
	for i := range beta {
		beta[i] = values[i] / matrix[i][i]
	}

	alpha := make([][]float64, dim)
	for i := range alpha {
		alpha[i] = make([]float64, dim)

		for j := range alpha[i] {
			alpha[i][j] = -matrix[i][j] / matrix[i][i]
			if i == j {
				alpha[i][j] = 0
			}
		}
	}
	return alpha, beta
}

func SimpleIteration(matrix [][]float64, values []float64, targetPrecision float64, useSeidel bool) ([]float64, int64, error) {
	dim := len(matrix)
	if dim == 0 {
		return nil, 0, fmt.Errorf("matrix is empty")
	}
	if targetPrecision <= 0 {
		return nil, 0, fmt.Errorf("target precision must be positive")
	}

	// TODO: if any matrix[i][i] == 0 try to swap rows
	for i := 0; i < dim; i++ {
		if matrix[i][i] == 0 {
			return nil, 0, fmt.Errorf("zero on main diagonal")
		}
	}

	// Making beta and alpha matrices
	alpha, beta := calcAlphaBeta(matrix, values)

	// The alpha norm calculation
	alphaNorm := 0.0
	for _, row := range alpha {
		sum := 0.0
		for _, elem := range row {
			sum += math.Abs(elem)
		}
		alphaNorm = math.Max(alphaNorm, sum)
	}

	for i := range alpha {
		alpha[i][i] = 0
	}

	// Iterations

	curIterX, prevIterX := make([]float64, dim), make([]float64, dim)
	// Iteration zero
	for i := range curIterX {
		curIterX[i] = beta[i]
	}

	// Perform iterations
	iterCount := int64(0)
	for {
		iterCount++
		for i := range prevIterX {
			prevIterX[i] = curIterX[i]
		}
		currentPrecision := math.Inf(-1)
		if useSeidel {
			curIterX = calcSeidelCurIterX(alpha, beta, prevIterX)
		} else {
			curIterX = calcCurIterX(alpha, beta, prevIterX)
		}

		// norm of x
		for i := range curIterX {
			currentPrecision = math.Max(currentPrecision, math.Abs(curIterX[i]-prevIterX[i]))
		}
		if alphaNorm < 1 {
			currentPrecision *= alphaNorm * (1 - alphaNorm)
		}

		if currentPrecision <= targetPrecision {
			break
		}
	}
	return curIterX, iterCount, nil
}

func SimpleIterationFromFile(filename string, precision float64) ([]float64, int64, error) {
	matrix, values, err := ReadMatrix(filename)
	if err != nil {
		return nil, 0, err
	}
	return SimpleIteration(matrix, values, precision, false)
}

func SeidelFromFile(filename string, precision float64) ([]float64, int64, error) {
	matrix, values, err := ReadMatrix(filename)
	if err != nil {
		return nil, 0, err
	}
	return SimpleIteration(matrix, values, precision, true)
}

//  endregion Iterations
