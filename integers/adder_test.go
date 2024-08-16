package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("adds two integers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertAddFunction(t, sum, expected)
	})
	t.Run("Should fail if expected is not 1", func(t *testing.T) {
		sum := Add(-1, 2)
		expected := 1
		assertAddFunction(t, sum, expected)
	})
}

func TestSubber(t *testing.T) {
	t.Run("subtracts two integers", func(t *testing.T) {
		sum := Subtract(4,2)
		expected := 2
		assertAddFunction(t, sum, expected)
	})
}

func assertAddFunction(t testing.TB, sum, expected int) {
	t.Helper()
	if expected != sum {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1,5)
	fmt.Println(sum)
	// Output: 6
}