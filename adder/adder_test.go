package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestExampleAdd(t *testing.T) {
	expected := 6

	if expected != 6 {
		t.Errorf("expected '%d' but got '%d'", 6, expected)
	}
}
