package mylib

import "testing"

// ex) $ go test -run TestAverage2
// ex) $ go test -v
// ex) $ go test
func TestAverage2(t *testing.T) {
	v := Average2([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, got", v)
	}
}
