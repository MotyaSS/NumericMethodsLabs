package lab3_2

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type SplineSegment struct {
	A, B, C, D, XStart, XEnd float64
}

func (s SplineSegment) GetPolynomial() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("S(x) = %.5f", s.A))

	if s.B >= 0 {
		sb.WriteString(fmt.Sprintf(" + %.5f(x - %.5f)", s.B, s.XStart))
	} else {
		sb.WriteString(fmt.Sprintf(" - %.5f(x - %.5f)", math.Abs(s.B), s.XStart))
	}

	if s.C >= 0 {
		sb.WriteString(fmt.Sprintf(" + %.5f(x - %.5f)^2", s.C, s.XStart))
	} else {
		sb.WriteString(fmt.Sprintf(" - %.5f(x - %.5f)^2", math.Abs(s.C), s.XStart))
	}

	if s.D >= 0 {
		sb.WriteString(fmt.Sprintf(" + %.5f(x - %.5f)^3", s.D, s.XStart))
	} else {
		sb.WriteString(fmt.Sprintf(" - %.5f(x - %.5f)^3", math.Abs(s.D), s.XStart))
	}

	return sb.String()
}

type CubicSpline struct {
	xValues, yValues, h, alpha, l, mu, z, a, b, c, d []float64
	Segments                                         []SplineSegment
}

func NewCubicSpline(xValues, yValues []float64) *CubicSpline {
	n := len(xValues)
	sortedIndices := make([]int, n)
	for i := range sortedIndices {
		sortedIndices[i] = i
	}
	sort.Slice(sortedIndices, func(i, j int) bool {
		return xValues[sortedIndices[i]] < xValues[sortedIndices[j]]
	})

	sortedX := make([]float64, n)
	sortedY := make([]float64, n)
	for i, idx := range sortedIndices {
		sortedX[i] = xValues[idx]
		sortedY[i] = yValues[idx]
	}

	spline := &CubicSpline{
		xValues:  sortedX,
		yValues:  sortedY,
		h:        make([]float64, n-1),
		alpha:    make([]float64, n-1),
		l:        make([]float64, n),
		mu:       make([]float64, n),
		z:        make([]float64, n),
		a:        make([]float64, n),
		b:        make([]float64, n-1),
		c:        make([]float64, n),
		d:        make([]float64, n-1),
		Segments: make([]SplineSegment, 0, n-1),
	}

	spline.initialize()
	spline.computeSplineCoefficients()
	spline.constructSegments()
	return spline
}

func (s *CubicSpline) initialize() {
	for i := 0; i < len(s.h); i++ {
		s.h[i] = s.xValues[i+1] - s.xValues[i]
		if s.h[i] <= 0 {
			panic("xValues must be in strictly increasing order.")
		}
	}

	for i := 1; i < len(s.alpha); i++ {
		s.alpha[i] = (3.0/s.h[i])*(s.yValues[i+1]-s.yValues[i]) - (3.0/s.h[i-1])*(s.yValues[i]-s.yValues[i-1])
	}
}

func (s *CubicSpline) computeSplineCoefficients() {
	s.l[0] = 1.0
	s.mu[0] = 0.0
	s.z[0] = 0.0

	for i := 1; i < len(s.h); i++ {
		s.l[i] = 2.0*(s.xValues[i+1]-s.xValues[i-1]) - s.h[i-1]*s.mu[i-1]
		if s.l[i] == 0 {
			panic("Singular matrix encountered in spline computation.")
		}
		s.mu[i] = s.h[i] / s.l[i]
		s.z[i] = (s.alpha[i] - s.h[i-1]*s.z[i-1]) / s.l[i]
	}

	s.l[len(s.l)-1] = 1.0
	s.z[len(s.z)-1] = 0.0
	s.c[len(s.c)-1] = 0.0

	for j := len(s.h) - 1; j >= 0; j-- {
		s.c[j] = s.z[j] - s.mu[j]*s.c[j+1]
		s.b[j] = (s.yValues[j+1]-s.yValues[j])/s.h[j] - s.h[j]*(s.c[j+1]+2.0*s.c[j])/3.0
		s.d[j] = (s.c[j+1] - s.c[j]) / (3.0 * s.h[j])
		s.a[j] = s.yValues[j]
	}
}

func (s *CubicSpline) constructSegments() {
	for i := 0; i < len(s.h); i++ {
		s.Segments = append(s.Segments, SplineSegment{
			A:      s.a[i],
			B:      s.b[i],
			C:      s.c[i],
			D:      s.d[i],
			XStart: s.xValues[i],
			XEnd:   s.xValues[i+1],
		})
	}
}

func (s *CubicSpline) Interpolate(x float64) float64 {
	if x < s.xValues[0] || x > s.xValues[len(s.xValues)-1] {
		panic("x is outside the interpolation range.")
	}

	interval := sort.Search(len(s.xValues), func(i int) bool {
		return s.xValues[i] >= x
	})
	if interval == len(s.xValues) || s.xValues[interval] != x {
		interval--
	}

	segment := s.Segments[interval]
	dx := x - segment.XStart
	return segment.A + segment.B*dx + segment.C*math.Pow(dx, 2) + segment.D*math.Pow(dx, 3)
}

func (s *CubicSpline) GetAllPolynomials() []string {
	polynomials := make([]string, len(s.Segments))
	for i, seg := range s.Segments {
		polynomials[i] = fmt.Sprintf("Interval [%.5f, %.5f]: %s", seg.XStart, seg.XEnd, seg.GetPolynomial())
	}
	return polynomials
}
