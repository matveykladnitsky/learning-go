package maxelement

import "testing"

func TestMaxPositiveNumbers(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected := 5
	if max(nums) != expected {
		t.Errorf("Ожидается %d, но получено %d", expected, max(nums))
	}
}

func TestMaxNegativeNumbers(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	expected := -1
	if max(nums) != expected {
		t.Errorf("Ожидается %d, но получено %d", expected, max(nums))
	}
}

func TestMaxMixedNumbers(t *testing.T) {
	nums := []int{-10, 0, 10, 5}
	expected := 10
	if max(nums) != expected {
		t.Errorf("Ожидается %d, но получено %d", expected, max(nums))
	}
}

func TestMaxEmptyArray(t *testing.T) {
	nums := []int{}
	expected := 0
	if max(nums) != expected {
		t.Errorf("Ожидается %d для пустого массива, но получено %d", expected, max(nums))
	}
}