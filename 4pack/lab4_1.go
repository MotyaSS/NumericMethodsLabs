package lab4_1

import (
	"fmt"
	"math"
	"strings"
)

func function(x, y, z float64) float64 {
	return -4*x*z + (3-4*x*x)*y + math.Exp(x*x)
}

func solution(x float64) float64 {
	return (math.Exp(x) + math.Exp(-x) - 1) * math.Exp(x*x)
}

func EulerMethodRungeRomberg(xValues, yValues, zValues []float64, h float64) {
	fmt.Println("Euler Method")
	fmt.Printf("%-5s %-10s %-15s %-15s %-15s %-15s %-15s\n", "Step", "x", "y_h", "y_h/2", "Error", "z", "result")
	fmt.Println(strings.Repeat("-", 95))

	n := len(xValues)
	xValuesHalfStep := make([]float64, 2*n-1)
	yValuesHalfStep := make([]float64, 2*n-1)
	zValuesHalfStep := make([]float64, 2*n-1)

	xValuesHalfStep[0] = xValues[0]
	yValuesHalfStep[0] = yValues[0]
	zValuesHalfStep[0] = zValues[0]

	fmt.Printf("%-5d %-10.2f %-15.6f %-15s %-15s %-15.6f %-15.6f\n",
		0, xValues[0], yValues[0], "-", "-", zValues[0], solution(xValues[0]))

	for i := 1; i < len(xValues); i++ {
		xPrev := xValues[i-1]
		yPrev := yValues[i-1]
		zPrev := zValues[i-1]

		yValues[i] = yPrev + h*zPrev
		zValues[i] = zPrev + h*function(xPrev, yPrev, zPrev)
		xValues[i] = xPrev + h
	}

	halfStep := h / 2.0
	for i := 1; i < len(xValuesHalfStep); i++ {
		xPrev := xValuesHalfStep[i-1]
		yPrev := yValuesHalfStep[i-1]
		zPrev := zValuesHalfStep[i-1]

		yValuesHalfStep[i] = yPrev + halfStep*zPrev
		zValuesHalfStep[i] = zPrev + halfStep*function(xPrev, yPrev, zPrev)
		xValuesHalfStep[i] = xPrev + halfStep
	}

	for i := 1; i < len(xValues); i++ {
		yH := yValues[i]
		yH2 := yValuesHalfStep[2*i]
		err := math.Abs((yH2 - yH) / (math.Pow(2, 2) - 1))

		fmt.Printf("%-5d %-10.2f %-15.6f %-15.6f %-15.6g %-15.6f %-15.6f\n",
			i, xValues[i], yH, yH2, err, zValues[i], solution(xValues[i]))
	}
}

func rungeKuttaStep(x, y, z, h float64) (float64, float64, float64) {
	k1_y := h * z
	L1_z := h * function(x, y, z)

	k2_y := h * (z + 0.5*L1_z)
	L2_z := h * function(x+0.5*h, y+0.5*k1_y, z+0.5*L1_z)

	k3_y := h * (z + 0.5*L2_z)
	L3_z := h * function(x+0.5*h, y+0.5*k2_y, z+0.5*L2_z)

	k4_y := h * (z + L3_z)
	L4_z := h * function(x+h, y+k3_y, z+L3_z)

	y_next := y + (k1_y+2*k2_y+2*k3_y+k4_y)/6.0
	z_next := z + (L1_z+2*L2_z+2*L3_z+L4_z)/6.0
	x_next := x + h

	return x_next, y_next, z_next
}

func RungeKuttaMethod(xValues, yValues, zValues []float64, h float64) {
	fmt.Println("\nRunge Kutta Method")
	fmt.Printf("%-5s %-10s %-15s %-15s %-15s %-15s %-15s\n", "Step", "x", "y_h", "y_h/2", "Error", "z", "result")
	fmt.Println(strings.Repeat("-", 100))

	fmt.Printf("%-5d %-10.2f %-15.6f %-15s %-15s %-15.6f %-15.6f\n",
		0, xValues[0], yValues[0], "-", "-", zValues[0], solution(xValues[0]))

	for i := 1; i < len(xValues); i++ {
		x := xValues[i-1]
		y := yValues[i-1]
		z := zValues[i-1]

		k1_y := h * z
		L1_z := h * function(x, y, z)

		k2_y := h * (z + 0.5*L1_z)
		L2_z := h * function(x+0.5*h, y+0.5*k1_y, z+0.5*L1_z)

		k3_y := h * (z + 0.5*L2_z)
		L3_z := h * function(x+0.5*h, y+0.5*k2_y, z+0.5*L2_z)

		k4_y := h * (z + L3_z)
		L4_z := h * function(x+h, y+k3_y, z+L3_z)

		y_h := y + (k1_y+2*k2_y+2*k3_y+k4_y)/6.0
		z_h := z + (L1_z+2*L2_z+2*L3_z+L4_z)/6.0
		x_h := x + h

		h2 := h / 2.0
		y_halfStep1 := y
		z_halfStep1 := z

		k1_y = h2 * z_halfStep1
		L1_z = h2 * function(x, y_halfStep1, z_halfStep1)

		k2_y = h2 * (z_halfStep1 + 0.5*L1_z)
		L2_z = h2 * function(x+0.5*h2, y_halfStep1+0.5*k1_y, z_halfStep1+0.5*L1_z)

		k3_y = h2 * (z_halfStep1 + 0.5*L2_z)
		L3_z = h2 * function(x+0.5*h2, y_halfStep1+0.5*k2_y, z_halfStep1+0.5*L2_z)

		k4_y = h2 * (z_halfStep1 + L3_z)
		L4_z = h2 * function(x+h2, y_halfStep1+k3_y, z_halfStep1+L3_z)

		y_halfStep2 := y_halfStep1 + (k1_y+2*k2_y+2*k3_y+k4_y)/6.0
		z_halfStep2 := z_halfStep1 + (L1_z+2*L2_z+2*L3_z+L4_z)/6.0

		x += h2
		k1_y = h2 * z_halfStep2
		L1_z = h2 * function(x, y_halfStep2, z_halfStep2)

		k2_y = h2 * (z_halfStep2 + 0.5*L1_z)
		L2_z = h2 * function(x+0.5*h2, y_halfStep2+0.5*k1_y, z_halfStep2+0.5*L1_z)

		k3_y = h2 * (z_halfStep2 + 0.5*L2_z)
		L3_z = h2 * function(x+0.5*h2, y_halfStep2+0.5*k2_y, z_halfStep2+0.5*L2_z)

		k4_y = h2 * (z_halfStep2 + L3_z)
		L4_z = h2 * function(x+h2, y_halfStep2+k3_y, z_halfStep2+L3_z)

		y_h2 := y_halfStep2 + (k1_y+2*k2_y+2*k3_y+k4_y)/6.0

		err := math.Abs(y_h-y_h2) / (math.Pow(2, 4) - 1)

		yValues[i] = y_h
		zValues[i] = z_h
		xValues[i] = x_h

		fmt.Printf("%-5d %-10.2f %-15.6f %-15.6f %-15.6g %-15.6f %-15.6f\n",
			i, xValues[i], y_h, y_h2, err, zValues[i], solution(xValues[i]))
	}
}

func AdamsMethod(xValues, yValues, zValues []float64, h float64) {
	if len(xValues) < 4 {
		panic("The array size should be at least 4 for the Adams method")
	}

	fmt.Println("Adams Method")
	fmt.Printf("%-5s %-10s %-15s %-20s %-15s %-15s %-15s\n",
		"Step", "x", "y_h", "y_h (corrected)", "Error", "z", "result")
	fmt.Println(strings.Repeat("-", 95))

	fmt.Printf("%-5d %-10.2f %-15.6f %-20s %-15s %-15.6f %-15.6f\n",
		0, xValues[0], yValues[0], "-", "-", zValues[0], solution(xValues[0]))

	for i := 1; i < 4; i++ {
		xValues[i], yValues[i], zValues[i] = rungeKuttaStep(xValues[i-1], yValues[i-1], zValues[i-1], h)
		fmt.Printf("%-5d %-10.2f %-15.6f %-20s %-15s %-15.6f %-15s\n",
			i, xValues[i], yValues[i], "-", "-", zValues[i], "-")
	}

	for i := 4; i < len(xValues); i++ {
		y_h := yValues[i-1] + h*(55*zValues[i-1]-59*zValues[i-2]+
			37*zValues[i-3]-9*zValues[i-4])/24.0

		z_h := zValues[i-1] + h*(55*function(xValues[i-1], yValues[i-1], zValues[i-1])-
			59*function(xValues[i-2], yValues[i-2], zValues[i-2])+
			37*function(xValues[i-3], yValues[i-3], zValues[i-3])-
			9*function(xValues[i-4], yValues[i-4], zValues[i-4]))/24.0

		y_h_corrected := yValues[i-1] + h*(9*z_h+19*zValues[i-1]-
			5*zValues[i-2]+zValues[i-3])/24.0

		z_h_corrected := zValues[i-1] + h*(9*function(xValues[i-1]+h, y_h, z_h)+
			19*function(xValues[i-1], yValues[i-1], zValues[i-1])-
			5*function(xValues[i-2], yValues[i-2], zValues[i-2])+
			function(xValues[i-3], yValues[i-3], zValues[i-3]))/24.0

		err := math.Abs(y_h_corrected-y_h) / (math.Pow(2, 4) - 1)

		yValues[i] = y_h_corrected
		zValues[i] = z_h_corrected
		xValues[i] = xValues[i-1] + h

		fmt.Printf("%-5d %-10.2f %-15.6f %-20.6f %-15.6g %-15.6f %-15.6f\n",
			i, xValues[i], y_h, y_h_corrected, err, zValues[i], solution(xValues[i]))
	}
}

func Example() {
	x0 := 0.0
	xEnd := 1.0
	y0 := 1.0
	z0 := 0.0
	h := 0.1
	n := int((xEnd-x0)/h) + 1

	xEuler := make([]float64, n)
	yEuler := make([]float64, n)
	zEuler := make([]float64, n)
	xEuler[0] = x0
	yEuler[0] = y0
	zEuler[0] = z0

	xRunge := make([]float64, n)
	yRunge := make([]float64, n)
	zRunge := make([]float64, n)
	xRunge[0] = x0
	yRunge[0] = y0
	zRunge[0] = z0

	xAdams := make([]float64, n)
	yAdams := make([]float64, n)
	zAdams := make([]float64, n)
	xAdams[0] = x0
	yAdams[0] = y0
	zAdams[0] = z0

	EulerMethodRungeRomberg(xEuler, yEuler, zEuler, h)
	RungeKuttaMethod(xRunge, yRunge, zRunge, h)
	AdamsMethod(xAdams, yAdams, zAdams, h)
}
