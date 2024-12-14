package factorial

import "testing"

func TestFactorialZero(t *testing.T) {
	if factorial(0) != 1 {
		t.Errorf("Ожидается 1, но получено %d", factorial(0))
	}
}

func TestFactorialOne(t *testing.T) {
	if factorial(1) != 1 {
		t.Errorf("Ожидается 1, но получено %d", factorial(1))
	}
}

func TestFactorialFive(t *testing.T) {
	if factorial(5) != 120 {
		t.Errorf("Ожидается 120, но получено %d", factorial(5))
	}
}

func TestFactorialNegative(t *testing.T) {
	if factorial(-5) != 0 {
		t.Errorf("Ожидается 0 для отрицательных чисел, но получено %d", factorial(-5))
	}
}