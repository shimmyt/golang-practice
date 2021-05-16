package simplearraysum

import "testing"

func TestSimplearraysum(t *testing.T) {
	expected := simplearraysum([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if expected != 55 {
		t.Errorf("Comparison is incorrect, got %d, want %d", expected, 55)
	}
}
