package main

import "fmt"

func main() {
	var sum int32
	ar := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ar:", ar)

	for i := 0; i < len(ar); i++ {
		sum = sum + ar[i]
	}
	fmt.Println("Points:", sum)
}
