package equation_system_solve_algo

import (
	"math"
)

//x1 = sqrt(2lgx_2 + 1)

//x2 = (x_1^2 + a)/ax_1

func SimpleIteration(
	eq1, eq2, dEq1, dEq2 func(float64, float64) float64,
	firstX, firstY float64,
	eps float64,
) (float64, float64) {
	prevX1 := firstX
	prevX2 := firstY
	currX1 := eq1(prevX1, prevX2)
	currX2 := eq2(prevX1, prevX2)
	for {
		q := math.Max(dEq1(currX1, currX2), dEq2(currX1, currX2))
		if q/(1-q)*math.Max(math.Abs(currX1-prevX1), math.Abs(currX2-prevX2)) < eps {
			return currX1, currX2
		}
		prevX1 = currX1
		prevX2 = currX2
		currX1 = eq1(prevX1, prevX2)
		currX2 = eq2(prevX1, prevX2)
	}
}

func findDeterminant(a, b, c, d float64) float64 {
	return a*d - b*c
}

func Newton(
	eq1, eq2, dEq1X1, dEq1X2, dEq2X1, dEq2X2 func(float64, float64) float64,
	xFirst, yFirst float64,
	eps float64,
) (float64, float64) {
	prevX := xFirst
	prevY := yFirst

	for {
		dF1x1 := dEq1X1(prevX, prevY)
		dF1x2 := dEq1X2(prevX, prevY)
		dF2x1 := dEq2X1(prevX, prevY)
		dF2x2 := dEq2X2(prevX, prevY)

		f1 := eq1(prevX, prevY)
		f2 := eq2(prevX, prevY)

		J := findDeterminant(dF1x1, dF1x2, dF2x1, dF2x2)
		A1 := findDeterminant(f1, dF1x2, f2, dF2x2)
		A2 := findDeterminant(dF1x1, f1, dF2x1, f2)
		curX := prevX - A1/J
		curY := prevY - A2/J
		if math.Max(math.Abs(curX-prevX), math.Abs(curY-prevY)) < eps {
			return curX, curY
		}
		prevX, prevY = curX, curY
	}
}
