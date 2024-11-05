package rotation_algo

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadMatrix(filename string) (matrix [][]float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return make([][]float64, 0), err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Split(line, ",")
		var row []float64
		for _, numStr := range numStrs {
			numStr = strings.TrimSpace(numStr)
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return make([][]float64, 0), fmt.Errorf("invalid number format: %s", numStr)
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}
	if err := scanner.Err(); err != nil {
		return make([][]float64, 0), err
	}
	return matrix, nil
}

func RotationsFromFile(filename string, targetPrecision float64) (eigenvalues []float64, eigenvectors [][]float64, err error) {
	if targetPrecision <= 0 {
		return nil, nil, fmt.Errorf("precision must be above zero")
	}
	matrix, err := ReadMatrix(filename)
	if err != nil {
		return nil, nil, err
	}
	return Rotations(matrix, targetPrecision)
}

func Rotations(matrix [][]float64, targetPrecision float64) (eigenvalues []float64, eigenvectors [][]float64, err error) {
	aMatrix := CopyMatrix(matrix)
	eigenvectors = getIdentityMatrix(len(aMatrix))
	curPrecision := math.Inf(1)
	for curPrecision > targetPrecision {
		row, clm := findMaxAbsNonDiagonalPos(aMatrix)
		rotationMatrix := getRotationMatrix(aMatrix, row, clm)
		aMatrix = multiplyMatrices(multiplyMatrices(rotationMatrix, aMatrix), transposeMatrix(rotationMatrix))
		eigenvectors = multiplyMatrices(eigenvectors, rotationMatrix)
		// Calculate precision
		curPrecision = 0
		for i := range aMatrix {
			for j := range aMatrix[i] {
				if i != j {
					curPrecision += math.Pow(aMatrix[i][j], 2)
				}
			}
		}
	}
	eigenvalues = make([]float64, len(aMatrix))
	for i := range aMatrix {
		eigenvalues[i] = aMatrix[i][i]
	}
	return
}

func CopyMatrix(matrix [][]float64) (res [][]float64) {
	res = make([][]float64, len(matrix))
	for i := range matrix {
		res[i] = make([]float64, len(matrix[i]))
		copy(res[i], matrix[i])
	}
	return
}

func findMaxAbsNonDiagonalPos(matrix [][]float64) (row int, column int) {
	rowM, clmM := 0, 1
	for i := range matrix {
		for j := i + 1; j < len(matrix[i]); j++ {
			if math.Abs(matrix[i][j]) > math.Abs(matrix[rowM][clmM]) {
				rowM = i
				clmM = j
			}
		}
	}
	return rowM, clmM
}

func getRotationMatrix(matrix [][]float64, row, clm int) [][]float64 {
	res := getIdentityMatrix(len(matrix))
	var angle float64
	if matrix[row][row] == matrix[clm][clm] {
		angle = math.Pi / 4
	} else {
		angle = 1. / 2 * math.Atan((2*matrix[row][clm])/(matrix[row][row]-matrix[clm][clm]))
	}
	sin, cos := math.Sin(angle), math.Cos(angle)
	res[row][row] = cos
	res[row][clm] = -sin
	res[clm][row] = sin
	res[clm][clm] = cos
	return res
}

func getIdentityMatrix(dim int) [][]float64 {
	res := make([][]float64, dim)
	for i := range res {
		res[i] = make([]float64, dim)
		for j := range res[i] {
			res[i][j] = 0
			if i == j {
				res[i][j] = 1
			}
		}
	}
	return res
}

func multiplyMatrices(a, b [][]float64) (res [][]float64) {
	res = make([][]float64, len(a))
	for i := range a {
		res[i] = make([]float64, len(b[0]))
		for j := range b[0] {
			for k := range a[0] {
				res[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return
}

func transposeMatrix(matrix [][]float64) (res [][]float64) {
	res = make([][]float64, len(matrix[0]))
	for i := range res {
		res[i] = make([]float64, len(matrix))
		for j := range matrix {
			res[i][j] = matrix[j][i]
		}
	}
	return
}
