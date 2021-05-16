package comparetriplets

import "testing"

func TestComparetriplets(t *testing.T) {
	expected := comparetriplets([]int32{3, 2, 1}, []int32{1, 2, 3})
	if expected != [2]int32{1, 1} {
		t.Errorf("Comparison is incorrect, got %d, want %d", expected, [2]int32{1, 1})
	}
}
