package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		// print integers with %d
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Code example are a nice way to show the usage and document the code.
// Because they get compiled and tested, you can be assured that the code is correct.
// The code will always be complied but by adding the output comment, it will also be executed
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
