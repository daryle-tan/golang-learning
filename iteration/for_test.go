package iteration

import (
	// "fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("should repeat a character 5 times", func(t *testing.T) {
		repeated := Iterator("d")
		expected := "ddddd"

		if repeated != expected {
			t.Errorf("expected %q got %q", expected, repeated)
		}
	})
}