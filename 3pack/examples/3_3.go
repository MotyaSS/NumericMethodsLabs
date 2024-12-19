package examples

import (
	"fmt"

	lab3_3 "Chislaki/3pack/3_3"
)

func LeastSquaresExample() {
	xValues := []float64{-5.0, -3.0, -1.0, 1.0, 3.0, 5.0}
	yValues := []float64{-1.3734, -1.249, -0.7854, 0.7854, 1.249, 1.3734}

	resultX1 := lab3_3.LSMethodFirstPower(xValues, yValues)
	fmt.Println("Solving a system of equations:")
	for _, v := range resultX1 {
		fmt.Println(v)
	}

	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("A polynomial of the 1st degree: %fx + %f\n", resultX1[1], resultX1[0])
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("Sum of squared errors: %f\n", lab3_3.SumSquaredErrorsFirst(resultX1, xValues, yValues))
	fmt.Println("----------------------------------------------------------------------------------")

	resultX2 := lab3_3.LSMethodSecondPower(xValues, yValues)
	fmt.Println("Solving a system of equations:")
	for _, v := range resultX2 {
		fmt.Println(v)
	}
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("A polynomial of the 2nd degree: %fx^2 + %fx + %f\n", resultX2[2], resultX2[1], resultX2[0])
	fmt.Println("----------------------------------------------------------------------------------")
	fmt.Printf("Sum of squared errors: %f\n", lab3_3.SumSquaredErrorsSecond(resultX2, xValues, yValues))
	fmt.Println("----------------------------------------------------------------------------------")
}
