package main

import "math"

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) (float64, float64) {
	var result1 = (-b + math.Sqrt(math.Pow(b, 2)-4*a*b)) / 2 * a
	var result2 = (-b - math.Sqrt(math.Pow(b, 2)-4*a*b)) / 2 * a
	return result1, result2
}

func main() {
	// TODO: call the function computeQuadraticFormula
	computeQuadraticFormula(2, 4, 7)
}
