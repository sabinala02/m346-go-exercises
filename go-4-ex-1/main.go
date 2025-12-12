package main

// TODO: implement the function computeGrade
func computeGrade(points float64, maxPoints float64) float64 {
	return (points*5)/maxPoints + 1
}

func main() {
	// TODO: call the function computeGrade
	computeGrade(12, 18)
}
