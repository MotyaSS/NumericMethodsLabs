package lab3_1

func LagrangePolynomial(funcToInterpolate func(float64) float64, xPoints []float64) []float64 {
	n := len(xPoints)
	L := make([]float64, n)
	for i := 0; i < n; i++ {
		L[i] = funcToInterpolate(xPoints[i])
		for j := 0; j < n; j++ {
			if i != j {
				L[i] /= xPoints[i] - xPoints[j]
			}
		}
	}
	return L
}

func separatedDifferences(funcEval func(float64) float64, x []float64) []float64 {
	size := (1 + len(x)) * len(x) / 2
	f := make([]float64, size)
	for i := range x {
		f[i] = funcEval(x[i])
	}
	k := len(x)
	p := 0
	e := len(x)
	for j := 1; j < len(x); j++ {
		for i := 0; i < len(x)-j; i++ {
			f[e] = (f[p+i] - f[p+i+1]) / (x[i] - x[i+j])
			e++
		}
		p += k
		k--
	}
	return f
}

func NewtonPolynomial(funcEval func(float64) float64, x []float64) []float64 {
	P := make([]float64, len(x))
	f := separatedDifferences(funcEval, x)
	p := 0
	k := len(x)
	for i := range P {
		P[i] = f[p]
		p += k
		k--
	}
	return P
}
