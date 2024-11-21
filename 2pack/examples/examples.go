package examples

import (
	"fmt"
	"math"

	"Chislaki/2pack/equation_solve_algo"
	"Chislaki/2pack/equation_system_solve_algo"
)

func equationXEquals(x float64) float64 {
	return math.Sqrt((math.Sin(x) + .5) / 2)
}

func equationDerivative(x float64) float64 {
	return math.Cos(x) - 4*x
}

func equation(x float64) float64 {
	return math.Sin(x) - 2*x*x + .5
}

func ExampleEquationSolve() {
	left, right := 0.5, 1.
	precision := .000001
	x, err := equation_solve_algo.SimpleIteration(equationXEquals, left, right, precision)
	if err != nil {
		println(err)
		return
	}
	fmt.Println("Простые итерации\t", x)

	x, err = equation_solve_algo.Newton(equation, equationDerivative, left, precision)

	if err != nil {
		println(err)
		return
	}
	fmt.Println("Ньютон\t\t\t", x)

	x, err = equation_solve_algo.Dichotomy(equation, left, right, precision)
	if err != nil {
		println(err)
		return
	}
	fmt.Println("Дихотомия\t\t", x)

	x, err = equation_solve_algo.Secant(equation, left, right, precision)
	if err != nil {
		println(err)
		return
	}
	fmt.Println("Секущие\t\t\t", x)
}

// region system

// x1 = a + cosx2
// x2 = a + sinx1

const a = 1

func firstEquation(_, x2 float64) float64 {
	return a + math.Cos(x2)
}

func secondEquation(x1, _ float64) float64 {
	return a + math.Cos(x1)
}

func firstEqDifferentialX1(_, _ float64) float64 {
	return 0
}

func firstEqDifferentialX2(_, x2 float64) float64 {
	return -math.Sin(x2)
}

func secondEqDifferentialX1(x1, _ float64) float64 {
	return math.Cos(x1)
}

func secondEqDifferentialX2(_, _ float64) float64 {
	return 0
}

func ExampleSystemSolve() {
	prec := .0001
	fmt.Println(equation_system_solve_algo.SimpleIteration(
		firstEquation, secondEquation,
		firstEqDifferentialX2, secondEqDifferentialX1,
		1., 1.5,
		prec,
	))

	fmt.Println(equation_system_solve_algo.Newton(
		firstEquation, secondEquation,
		firstEqDifferentialX1, firstEqDifferentialX2,
		secondEqDifferentialX1, secondEqDifferentialX2,
		1., 1.5,
		prec,
	))
}

// endregion system
