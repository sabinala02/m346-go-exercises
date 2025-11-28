package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

	fibs[2] = fibs[0] + fibs[1]
	fibs[3] = 3
	fibs[4] = 5

	fibs = append(fibs, 13)
	fibs = append(fibs, 21)
	fibs = append(fibs, 34)

	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
