package lab3_3

import (
	"math"
)

// LSMethodFirstPower calculates coefficients for linear regression (first degree polynomial)
func LSMethodFirstPower(x, y []float64) []float64 {
	n := len(x)
	var sumX, sumY, sumX2, sumXY float64

	for i := 0; i < n; i++ {
		sumX += x[i]
		sumY += y[i]
		sumX2 += x[i] * x[i]
		sumXY += x[i] * y[i]
	}

	a1 := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumX2 - sumX*sumX)
	a0 := (sumY - a1*sumX) / float64(n)

	return []float64{a0, a1}
}

// SumSquaredErrorsFirst calculates sum of squared errors for first degree polynomial
func SumSquaredErrorsFirst(xResult, x, y []float64) float64 {
	var result float64
	functionF := make([]float64, len(x))

	for i := range x {
		functionF[i] = xResult[1]*x[i] + xResult[0]
	}

	for i := range x {
		result += math.Pow(functionF[i]-y[i], 2)
	}

	return result
}

// LSMethodSecondPower calculates coefficients for quadratic regression (second degree polynomial)
func LSMethodSecondPower(x, y []float64) []float64 {
	n := len(x)
	var sumX, sumY, sumX2, sumX3, sumX4, sumXY, sumX2Y float64

	for i := 0; i < n; i++ {
		xi := x[i]
		yi := y[i]
		xi2 := xi * xi
		xi3 := xi2 * xi
		xi4 := xi3 * xi

		sumX += xi
		sumY += yi
		sumX2 += xi2
		sumX3 += xi3
		sumX4 += xi4
		sumXY += xi * yi
		sumX2Y += xi2 * yi
	}

	matrixA := [][]float64{
		{float64(n), sumX, sumX2},
		{sumX, sumX2, sumX3},
		{sumX2, sumX3, sumX4},
	}

	vectorB := []float64{sumY, sumXY, sumX2Y}
	return solveLinearSystem(matrixA, vectorB)
}

// solveLinearSystem solves system of linear equations using Gaussian elimination
func solveLinearSystem(matrix [][]float64, vector []float64) []float64 {
	n := len(vector)
	result := make([]float64, n)
	augmentedMatrix := make([][]float64, n)

	// Create augmented matrix
	for i := 0; i < n; i++ {
		augmentedMatrix[i] = make([]float64, n+1)
		copy(augmentedMatrix[i], matrix[i])
		augmentedMatrix[i][n] = vector[i]
	}

	// Forward elimination
	for i := 0; i < n; i++ {
		// Partial pivoting
		maxRow := i
		for k := i + 1; k < n; k++ {
			if math.Abs(augmentedMatrix[k][i]) > math.Abs(augmentedMatrix[maxRow][i]) {
				maxRow = k
			}
		}

		if maxRow != i {
			augmentedMatrix[i], augmentedMatrix[maxRow] = augmentedMatrix[maxRow], augmentedMatrix[i]
		}

		// Elimination
		for k := i + 1; k < n; k++ {
			factor := augmentedMatrix[k][i] / augmentedMatrix[i][i]
			for j := i; j <= n; j++ {
				augmentedMatrix[k][j] -= factor * augmentedMatrix[i][j]
			}
		}
	}

	// Back substitution
	for i := n - 1; i >= 0; i-- {
		result[i] = augmentedMatrix[i][n]
		for j := i + 1; j < n; j++ {
			result[i] -= augmentedMatrix[i][j] * result[j]
		}
		result[i] /= augmentedMatrix[i][i]
	}

	return result
}

// SumSquaredErrorsSecond calculates sum of squared errors for second degree polynomial
func SumSquaredErrorsSecond(xResult, x, y []float64) float64 {
	var result float64
	functionF := make([]float64, len(x))

	for i := range x {
		functionF[i] = xResult[2]*x[i]*x[i] + xResult[1]*x[i] + xResult[0]
	}

	for i := range x {
		result += math.Pow(functionF[i]-y[i], 2)
	}

	return result
}
