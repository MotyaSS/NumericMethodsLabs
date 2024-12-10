package examples

import (
	"fmt"
	"math"

	"Chislaki/2pack/equation_solve_algo"
	essa "Chislaki/2pack/equation_system_solve_algo"
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

//
// x1 = cos(x2) + a
//

const a = 1

func firstEquation(x1, x2 float64) float64 {
	return math.Cos(x2) + a // x1 - cos(x2) - a = 0
}

//
// x2 = sin(x1) + a
//

func secondEquation(x1, x2 float64) float64 {
	return math.Sin(x1) + a // x2 - sin(x1) - a = 0
}

func equationsSystem(x *essa.Matrix) *essa.Matrix {
	f := essa.NewMatrix(2, 1)
	f.Set(0, 0, x.Get(0, 0)-math.Cos(x.Get(1, 0))-1)
	f.Set(1, 0, x.Get(1, 0)-math.Sin(x.Get(0, 0))-1)
	return f
}

func jacobiMatrix(x *essa.Matrix) *essa.Matrix {
	j := essa.NewMatrix(2, 2)
	j.Set(0, 0, 1)
	j.Set(0, 1, math.Sin(x.Get(1, 0)))
	j.Set(1, 0, math.Sin(x.Get(0, 0)))
	j.Set(1, 1, 1)
	return j
}

func ExampleSystemSolve() {
	const eps = .001
	r1, r2, err := essa.SimpleIterations(
		firstEquation, secondEquation,
		1., 1.5,
		eps,
	)
	fmt.Println("Простые итерации")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r1, r2)
	}

	x1 := essa.NewMatrix(2, 1)
	x1.Set(0, 0, 0.25)
	x1.Set(1, 0, 0.25)
	fmt.Println("Newton method")
	result1 := essa.NewtonMethod(equationsSystem, jacobiMatrix, x1, eps)
	fmt.Println(result1.Get(0, 0), result1.Get(1, 0))
}

// endregion system
