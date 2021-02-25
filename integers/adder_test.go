package integers

import "testing"
import (
	"fmt"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	expected := 4

	if sum := Add(2, 2); sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}