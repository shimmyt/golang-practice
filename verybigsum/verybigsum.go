package verybigsum

func verybigsum(ar []int64) int64 {
	var sum int64

	for i := 0; i < len(ar); i++ {
		sum = sum + ar[i]
	}

	return sum
}
