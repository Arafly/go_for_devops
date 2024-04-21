package iteration

import "testing"

func TestRepeat (t *testing.T) {
	repeated := Repeat ("c", 5)
	expected := "ccccc"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s' instead", expected, repeated)
	}
}

// BenchmarkRepeat is a function that benchmarks the Repeat function
// It runs the Repeat function 5 times and records the time it takes to run
// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("c", 5)
	}
}