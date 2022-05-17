package paradime

import (
	"testing"
)

func TestReverse(t *testing.T) {

	words := []string{"ape", "donkey", "cow"}
	paradimes := []string{"epa", "yeknod", "woc"}

	for i, word := range words {
		got := Reverse(word)
		expected := paradimes[i]

		if got != expected {
			t.Errorf("Expected '%s' but got '%s'", expected, got)
		}
	}

}
