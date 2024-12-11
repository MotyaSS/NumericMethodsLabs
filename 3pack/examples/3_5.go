package examples

import (
	"fmt"

	lab3_5 "Chislaki/3pack/3_5"
)

func ExampleIntegration() {
	h1 := 0.5
	h2 := 0.25
	x0 := 0.0
	x1 := 2.0

	integrator := lab3_5.Integrator{
		Function: func(x float64) float64 {
			return (x * x) / (x*x + 16)
		},
	}

	rectangleValue1 := integrator.RectangleMethod(x0, x1, h1)
	rectangleValue2 := integrator.RectangleMethod(x0, x1, h2)
	rectangleRunge := lab3_5.RungeRomberg(rectangleValue1, rectangleValue2, 1, h1, h2)

	fmt.Println("Rectangle Method")
	fmt.Println("--------------------------------------------")
	fmt.Printf("rectangleValue1 = %.5f\n", rectangleValue1)
	fmt.Printf("rectangleValue2 = %.5f\n", rectangleValue2)
	fmt.Printf("Runge Romberg = %.5f\n", rectangleRunge)
	fmt.Println("--------------------------------------------")
	fmt.Println()

	trapezoidValue1 := integrator.TrapezoidMethod(x0, x1, h1)
	trapezoidValue2 := integrator.TrapezoidMethod(x0, x1, h2)
	trapezoidRunge := lab3_5.RungeRomberg(trapezoidValue1, trapezoidValue2, 2, h1, h2)

	fmt.Println("Trapezoid Method")
	fmt.Println("--------------------------------------------")
	fmt.Printf("trapezoidValue1 = %.5f\n", trapezoidValue1)
	fmt.Printf("trapezoidValue2 = %.5f\n", trapezoidValue2)
	fmt.Printf("Runge Romberg = %.5f\n", trapezoidRunge)
	fmt.Println("--------------------------------------------")
	fmt.Println()

	simpsonValue1 := integrator.SimpsonMethod(x0, x1, h1)
	simpsonValue2 := integrator.SimpsonMethod(x0, x1, h2)
	simpsonRunge := lab3_5.RungeRomberg(simpsonValue1, simpsonValue2, 4, h1, h2)

	fmt.Println("Simpson Method")
	fmt.Println("--------------------------------------------")
	fmt.Printf("simpsonValue1 = %.5f\n", simpsonValue1)
	fmt.Printf("simpsonValue2 = %.5f\n", simpsonValue2)
	fmt.Printf("Runge Romberg = %.5f\n", simpsonRunge)
	fmt.Println("--------------------------------------------")
	fmt.Println()

	fmt.Println("First:")
	integrator.CreateTable(h1, x0, x1)

	fmt.Println("Second:")
	integrator.CreateTable(h2, x0, x1)
}
