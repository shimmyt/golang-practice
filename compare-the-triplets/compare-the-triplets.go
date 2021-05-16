package main

import "fmt"

func main() {
	var points [2]int32
	a := [3]int32{5, 6, 7}
	b := [3]int32{3, 6, 10}
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			points[0] = points[0] + 1
		} else if a[i] < b[i] {
			points[1] = points[1] + 1
		}
	}
	fmt.Println("Points:", points)
}
