package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("expected: '%d', result: '%d'", expected, result)
	}
}

func ExampleSum() {
	result := sum(1, 1)
	fmt.Println(result)
	// Output: 2
}