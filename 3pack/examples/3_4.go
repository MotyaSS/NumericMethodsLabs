package examples

import (
	"fmt"

	lab3_4 "Chislaki/3pack/3_4"
)

func DerivativesExample() {
	xValues := []float64{0.0, 0.5, 1.0, 1.5, 2.0}
	yValues := []float64{0.0, 0.97943, 1.8415, 2.4975, 2, 9093}
	xToEvaluate := 1.0

	var inputPoints []lab3_4.Point
	for i := 0; i < len(xValues); i++ {
		inputPoints = append(inputPoints, lab3_4.Point{X: xValues[i], Y: yValues[i]})
	}

	leftHandedDerivative, rightHandedDerivative, firstDer, secondDer := lab3_4.CalculateDerivatives(inputPoints, xToEvaluate)

	fmt.Printf("Left-handed derivative: %.5f\n", leftHandedDerivative)
	fmt.Printf("Right-handed derivative: %.5f\n", rightHandedDerivative)
	fmt.Printf("First derivative: %.5f\n", firstDer)
	fmt.Printf("Second derivative: %.5f\n", secondDer)

	fmt.Printf("Check: %.5f\n", (leftHandedDerivative+rightHandedDerivative)/2)
}
