package equation_system_solve_algo

import (
	"math"
)

// SimpleIterations solves system of two equations with Simple Iterations method
func SimpleIterations(f1, f2 func(float64, float64) float64, x10, x20 float64, epsilon float64) (float64, float64, error) {
	x1, x2 := x10, x20
	for {
		nextX1 := f1(x1, x2)
		nextX2 := f2(x1, x2)
		if math.Abs(nextX1-x1) < epsilon && math.Abs(nextX2-x2) < epsilon {
			return nextX1, nextX2, nil
		}
		x1, x2 = nextX1, nextX2
	}
}

func NewtonMethod(f func(*Matrix) *Matrix,
	j func(*Matrix) *Matrix,
	x *Matrix, eps float64) *Matrix {

	var prevX *Matrix
	for {
		prevX = x
		jacobian := j(x)
		fx := f(x)
		delta := jacobian.InverseMatrix().Multiply(fx)
		x = x.Subtract(delta)

		if x.Subtract(prevX).VecNormC() <= eps {
			break
		}
	}
	return x
}
