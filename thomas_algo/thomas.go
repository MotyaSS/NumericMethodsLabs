package thomas_algo

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func stringToIntSlice(str string, sep string) []int64 {
	resStr := strings.Split(str, sep)
	var res []int64

	for ind := range resStr {
		resStr[ind] = strings.ReplaceAll(resStr[ind], " ", "")
		val, err := strconv.Atoi(resStr[ind])
		if err == nil {
			res = append(res, int64(val))
		}
	}
	return res
}

func calculatePQ(a, b, c, d []int64, dim int) ([]big.Rat, []big.Rat) {
	p := make([]big.Rat, dim+1)
	q := make([]big.Rat, dim+1)
	p[1].SetFrac64(-c[1], b[1])
	q[1].SetFrac64(d[1], b[1])
	for i := 2; i <= dim; i++ {
		denom := big.NewRat(a[i], 1)
		denom.Mul(denom, &p[i-1]).Add(denom, big.NewRat(b[i], 1))

		p[i].SetFrac64(-c[i], 1).Quo(&p[i], denom)
		q[i].SetFrac64(d[i], 1).Sub(&q[i], big.NewRat(a[i], 1).Mul(&q[i-1], big.NewRat(a[i], 1))).Quo(&q[i], denom)
	}
	return p, q
}

func getDiagonals(scanner *bufio.Scanner) ([][]int64, int) {
	var diagonal = make([][]int64, 3)
	dimension := -1
	for i := 0; i <= 2 && scanner.Scan(); i++ {
		diagonal[i] = append(diagonal[i], stringToIntSlice(scanner.Text(), ",")...)
		if dimension == -1 {
			dimension = len(diagonal[i])
		}
		if dimension != len(diagonal[i]) {
			panic("well well well")
		}
		diagonal[i] = append(make([]int64, 1), diagonal[i]...)
	}
	return diagonal, dimension
}

func isStable(a, b, c []int64, dim int) bool {
	for i := 1; i < dim; i++ {
		if math.Abs(float64(b[i])) < math.Abs(float64(a[i]))+math.Abs(float64(c[i])) {
			return false
		}
	}
	return true
}
func ThomasAlgo(a, b, c, d []int64, dim int) ([]big.Rat, *big.Rat, error) {
	// Check if the matrix is stable
	if !isStable(a, b, c, dim) {
		return nil, nil, fmt.Errorf("the matrix is not stable")
	}
	p, q := calculatePQ(a, b, c, d, dim)

	//  find the solution x
	x := make([]big.Rat, dim+1)
	x[dim].Set(&q[dim])
	for i := dim - 1; i >= 1; i-- {
		x[i].Mul(&p[i], &x[i+1]).Add(&x[i], &q[i])
	}

	// Calculate the determinant
	det := big.NewRat(1, 1)
	for i := 1; i <= dim; i++ {
		det.Mul(det, big.NewRat(b[i], 1))
	}

	return x[1:], det, nil
}

func checkSolution(a, b, c, d []int64, x []big.Rat, dim int) bool {
	for i := 1; i < dim-1; i++ {
		if a[i]*x[i-1].Num().Int64()+b[i]*x[i].Num().Int64()+c[i]*x[i+1].Num().Int64() != d[i] {
			return false
		}
	}
	return true
}

func ThomasFile(filename string) ([]big.Rat, *big.Rat, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)

	// Get diagonals and dimension
	diagonal, dimension := getDiagonals(scanner)

	// Get the d vector
	d := make([]int64, 1)
	if scanner.Scan() {
		d = append(d, stringToIntSlice(scanner.Text(), ",")...)
	}

	// Use ThomasAlgo to find the solution and determinant
	x, det, err := ThomasAlgo(diagonal[0], diagonal[1], diagonal[2], d, dimension)
	if err != nil {
		return nil, nil, err
	}

	// Check the solution
	if !checkSolution(diagonal[0][1:], diagonal[1][1:], diagonal[2][1:], d[1:], x, dimension) {
		return nil, nil, fmt.Errorf("the solution is incorrect")
	}
	return x, det, nil
}
