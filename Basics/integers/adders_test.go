package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T){
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d' instead", expected, sum)
	}
}

// ExampleAdd is a function that demonstrates how to use the Add function
// It runs the Add function and prints the result
// go test -v -run=ExampleAdd
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}