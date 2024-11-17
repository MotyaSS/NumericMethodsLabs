package equation_solve_algo

import (
	"errors"
	"math"
)

func Newton(f, df func(float64) float64, x0, epsilon float64) (float64, error) {
	x := x0
	for {
		nextX := x - f(x)/df(x)
		if math.Abs(nextX-x) < epsilon {
			return nextX, nil
		}
		x = nextX
	}
}

func Dichotomy(f func(float64) float64, a, b, epsilon float64) (float64, error) {
	if a >= b {
		return 0, errors.New("invalid interval: a must be less than b")
	}

	for {
		x := (a + b) / 2
		if b-a < epsilon {
			return x, nil
		}
		if f(a)*f(x) < 0 {
			b = x
		} else {
			a = x
		}
	}
}

func Secant(f func(float64) float64, x0, x1, epsilon float64) (float64, error) {
	for {
		nextX := x1 - f(x1)*(x1-x0)/(f(x1)-f(x0))
		if math.Abs(nextX-x1) < epsilon {
			return nextX, nil
		}
		x0, x1 = x1, nextX
	}
}

func SimpleIteration(f func(float64) float64, a, b, epsilon float64) (float64, error) {
	if a >= b {
		return 0, errors.New("invalid interval: a must be less than b")
	}

	x := (a + b) / 2
	for {
		nextX := f(x)
		if math.Abs(nextX-x) < epsilon {
			return nextX, nil
		}
		x = nextX
	}
}
