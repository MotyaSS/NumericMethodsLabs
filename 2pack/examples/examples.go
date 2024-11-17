package examples

import (
	"fmt"
	"math"

	"Chislaki/2pack/equation_solve_algo"
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
