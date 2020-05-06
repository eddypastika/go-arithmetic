package go_arithmetic_test

import (
	arithmetic "github.com/eddypastika/go-arithmetic"
	"testing"
)

func TestMultiply(t *testing.T) {
	a, b := 10, 20
	res := arithmetic.Multiply(a, b)
	if res != 200 {
		t.Errorf("got %v, expected %v \n", res, 200)
	}
}

func TestAdd(t *testing.T) {
	a, b := 10, 20
	res := arithmetic.Add(a, b)
	if res != 30 {
		t.Errorf("got %v, expected %v \n", res, 30)
	}
}

func TestSubtract(t *testing.T) {
	a, b := 10, 20
	res := arithmetic.Subtract(a, b)
	if res != -10 {
		t.Errorf("got %v, expected %v \n", res, -10)
	}
}
