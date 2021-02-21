package calculator

import "testing"

func TestAdd(t *testing.T) {
	if got, want := Add(1, 2), 3; got != want {
		t.Errorf("add method produced wrong result. expected: %d, got: %d", want, got)
	}
}

func TestSub(t *testing.T) {
	if got, want := Subtract(4, 2), 2; got != want {
		t.Errorf("sub method produced wrong result. expected: %d, got: %d", want, got)
	}
}
