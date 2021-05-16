package comparetriplets

func comparetriplets(a []int32, b []int32) [2]int32 {
	var points [2]int32

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			points[0] = points[0] + 1
		} else if a[i] < b[i] {
			points[1] = points[1] + 1
		}
	}

	return points
}
