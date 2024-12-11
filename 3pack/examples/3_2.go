package examples

import (
	"fmt"

	lab3_2 "Chislaki/3pack/3_2"
)

func CubicSplineExample() {
	xValues := []float64{-3.0, -1.0, 1.0, 3.0, 5.0}
	yValues := []float64{-1.2490, -0.78540, 0.78540, 1.2490, 1.3734}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("An error occurred:", r)
		}
	}()

	spline := lab3_2.NewCubicSpline(xValues, yValues)

	fmt.Println("Cubic Spline Polynomials for Each Interval:")
	for _, poly := range spline.GetAllPolynomials() {
		fmt.Println(poly)
	}

	xToInterpolate := -.5
	interpolatedValue := spline.Interpolate(xToInterpolate)
	fmt.Printf("Interpolated value at x = %.5f is f(%.5f) = %.5f\n", xToInterpolate, xToInterpolate, interpolatedValue)
}
