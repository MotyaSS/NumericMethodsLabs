package lab3_5

import (
	"fmt"
	"math"
)

type Integrator struct {
	Function func(float64) float64
}

func (i *Integrator) SecondDerivativeFunction(x float64) float64 {
	denominator := x*x + 4
	return (8*x*x)/(denominator*denominator*denominator) - (2 / (denominator * denominator))
}

func (i *Integrator) ErrorEstimateRectangleMethod(h, a, b float64) float64 {
	M2 := math.Max(math.Abs(i.SecondDerivativeFunction(a)), math.Abs(i.SecondDerivativeFunction(b)))
	return (1.0 / 24.0) * math.Pow(h, 2) * M2 * (b - a)
}

func (i *Integrator) RectangleMethod(a, b, h float64) float64 {
	sum := 0.0
	n := int(math.Ceil((b - a) / h))

	for j := 0; j < n; j++ {
		xi := a + float64(j)*h
		midpoint := xi + h/2
		sum += i.Function(midpoint)
	}

	return sum * h
}

func (i *Integrator) ErrorEstimateTrapezoidMethod(h, a, b float64) float64 {
	M2 := math.Max(math.Abs(i.SecondDerivativeFunction(a)), math.Abs(i.SecondDerivativeFunction(b)))
	return ((b - a) / 12.0) * h * h * M2
}

func (i *Integrator) TrapezoidMethod(a, b, h float64) float64 {
	n := int(math.Ceil((b - a) / h))
	sum := 0.5 * (i.Function(a) + i.Function(b))

	for j := 1; j < n; j++ {
		xi := a + float64(j)*h
		sum += i.Function(xi)
	}

	return sum * h
}

func (i *Integrator) SimpsonMethod(a, b, h float64) float64 {
	n := int(math.Ceil((b - a) / h))
	if n%2 == 1 {
		n++
	}

	sum := i.Function(a) + i.Function(b)

	for j := 1; j < n; j++ {
		xi := a + float64(j)*h
		if j%2 == 0 {
			sum += 2 * i.Function(xi)
		} else {
			sum += 4 * i.Function(xi)
		}
	}

	return (h / 3.0) * sum
}

func (i *Integrator) FourthDerivativeFunction(x float64) float64 {
	denominator := math.Pow(x*x+4, 3)
	return (48 * (x*x - 4)) / denominator
}

func (i *Integrator) ErrorEstimateSimpsonMethod(h, a, b float64) float64 {
	M4 := math.Max(math.Abs(i.FourthDerivativeFunction(a)), math.Abs(i.FourthDerivativeFunction(b)))
	return ((b - a) / 180.0) * h * h * h * h * M4
}

func (i *Integrator) CreateTable(h, x0, x1 float64) {
	numberOfSteps := int(math.Ceil(math.Abs(x1-x0) / h))

	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println(" i |  xi  |  yi  |  The rectangle method  |  The trapezoid method  |  Simpson's method  ")
	fmt.Println("----------------------------------------------------------------------------------------")

	for j := 0; j <= numberOfSteps; j++ {
		xi := x0 + float64(j)*h
		if xi > x1 {
			break
		}

		yi := i.Function(xi)
		rectangleValue := i.RectangleMethod(x0, xi, h)
		trapezoidValue := i.TrapezoidMethod(x0, xi, h)
		simpsonValue := i.SimpsonMethod(x0, xi, h)

		fmt.Printf(" %2d | %5.2f | %5.2f | %20.5f | %20.5f | %20.5f \n", j, xi, yi, rectangleValue, trapezoidValue, simpsonValue)
	}

	fmt.Println("----------------------------------------------------------------------------------------")

	errorEstimate1 := i.ErrorEstimateRectangleMethod(h, x0, x1)
	fmt.Printf("Error estimate in rectangle method: %.5f\n", errorEstimate1)

	errorEstimate2 := i.ErrorEstimateTrapezoidMethod(h, x0, x1)
	fmt.Printf("Error estimate in trapezoid method: %.5f\n", errorEstimate2)

	errorEstimate3 := i.ErrorEstimateSimpsonMethod(h, x0, x1)
	fmt.Printf("Error estimate in Simpson's method: %.5f\n", errorEstimate3)
}

func RungeRomberg(I_h1, I_h2, p, h1, h2 float64) float64 {
	denominator := math.Pow(h1/h2, p) - 1.0
	if math.Abs(denominator) < 1e-12 {
		panic("Denominator is too small, possible division by zero.")
	}
	return (I_h2 - I_h1) / denominator
}
