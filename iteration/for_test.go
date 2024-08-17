package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("should repeat a character 5 times", func(t *testing.T) {
		repeated := Iterator("d", 5)
		expected := "ddddd"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iterator("a", 5)
	}
}

func ExampleIterator() {
	repeat := Iterator("x", 3)
	fmt.Println(repeat)
	// Output: xxx
}
