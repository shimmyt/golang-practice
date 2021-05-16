package verybigsum

import "testing"

func TestVerybigsum(t *testing.T) {
	expected := verybigsum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	if expected != 5000000015 {
		t.Errorf("Comparison is incorrect, got %d, want %d", expected, 5000000015)
	}
}
