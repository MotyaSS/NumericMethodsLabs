package lab3_4

type Point struct {
	X, Y float64
}

func CalculateDerivatives(points []Point, x float64) (float64, float64, float64, float64) {
	index := 0
	for i := 1; i < len(points)-2; i++ {
		if points[i].X <= x && x <= points[i+1].X {
			index = i
			break
		}
	}

	leftHandedDerivative := (points[index+1].Y - points[index].Y) / (points[index+1].X - points[index].X)
	rightHandedDerivative := (points[index+2].Y - points[index+1].Y) / (points[index+2].X - points[index+1].X)
	firstDer := leftHandedDerivative +
		((rightHandedDerivative-leftHandedDerivative)/(points[index+2].X-points[index].X))*
			(2.0*x-points[index].X-points[index+1].X)
	secondDer := 2.0 * (rightHandedDerivative - leftHandedDerivative) / (points[index+2].X - points[index].X)

	return leftHandedDerivative, rightHandedDerivative, firstDer, secondDer
}
