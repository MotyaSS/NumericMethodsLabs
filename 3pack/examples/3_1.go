package examples

import (
	"fmt"
	"math"

	lab3_1 "Chislaki/3pack/3_1"
)

func LagrangePolynomialResult(L []float64, x []float64, value float64) float64 {
	res := 0.0

	for i := 0; i < len(L); i++ {
		tmp := L[i]

		for j := 0; j < len(x); j++ {
			if i != j {
				tmp *= value - x[j]
			}
		}

		res += tmp
	}

	return res
}

func PrintLagrangePolynomial(L []float64, x []float64) {
	for i := 0; i < len(L); i++ {
		if i != 0 && L[i] >= 0 {
			fmt.Print("+")
		}
		fmt.Printf("%.5f", L[i])

		for j := 0; j < len(x); j++ {
			if i != j {
				fmt.Print("(x")
				if x[j] >= 0 {
					fmt.Print("+")
				}
				fmt.Print(x[j], ")")
			}
		}
	}

	fmt.Println()
}

func NewtonPolynomialResult(P, x []float64, value float64) float64 {
	res := 0.0
	for i := range P {
		tmp := P[i]
		for j := 0; j < i; j++ {
			tmp *= value - x[j]
		}
		res += tmp
	}
	return res
}

func PrintNewtonPolynomial(P, x []float64) {
	for i, pi := range P {
		if i != 0 && pi >= 0 {
			fmt.Print("+")
		}
		fmt.Printf("%.5f", pi)
		for j := 0; j < i; j++ {
			fmt.Print("(x")
			if x[j] >= 0 {
				fmt.Print("+")
			}
			fmt.Print(x[j], ")")
		}
	}

	fmt.Println()
}

func InterpolationExample() {
	const X = -.5
	a := []float64{-3, -1, 1, 3}
	b := []float64{-3, 0, 1, 3}
	f := math.Atan
	fRes := f(X)
	fmt.Println("Function result: ", fRes)
	fmt.Println()

	pol := lab3_1.NewtonPolynomial(f, a)
	PrintNewtonPolynomial(pol, a)
	res := NewtonPolynomialResult(pol, a, X)
	fmt.Println("Newton result for", a, "\t", res)
	fmt.Println("Error: ", math.Abs(res-fRes))

	pol = lab3_1.NewtonPolynomial(f, b)
	PrintNewtonPolynomial(pol, b)
	res = NewtonPolynomialResult(pol, b, X)
	fmt.Println("Newton result for", b, "\t", res)
	fmt.Println("Error: ", math.Abs(res-fRes))

	pol = lab3_1.LagrangePolynomial(f, a)
	PrintLagrangePolynomial(pol, a)
	res = LagrangePolynomialResult(pol, a, X)
	fmt.Println("Lagrange result for", a, "", res)
	fmt.Println("Error: ", math.Abs(res-fRes))

	pol = lab3_1.LagrangePolynomial(f, b)
	PrintLagrangePolynomial(pol, b)
	res = LagrangePolynomialResult(pol, b, X)
	fmt.Println("Lagrange result for", b, "\t", res)
	fmt.Println("Error: ", math.Abs(res-fRes))
}
