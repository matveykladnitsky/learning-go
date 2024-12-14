package reverse

import "testing"

func TestReverseNormalString(t *testing.T) {
	input := "hello"
	expected := "olleh"
	if reverseString(input) != expected {
		t.Errorf("Ожидается %q, но получено %q", expected, reverseString(input))
	}
}

func TestReverseEmptyString(t *testing.T) {
	input := ""
	expected := ""
	if reverseString(input) != expected {
		t.Errorf("Ожидается %q, но получено %q", expected, reverseString(input))
	}
}

func TestReverseSingleChar(t *testing.T) {
	input := "a"
	expected := "a"
	if reverseString(input) != expected {
		t.Errorf("Ожидается %q, но получено %q", expected, reverseString(input))
	}
}

func TestReversePalindrome(t *testing.T) {
	input := "madam"
	expected := "madam"
	if reverseString(input) != expected {
		t.Errorf("Ожидается %q, но получено %q", expected, reverseString(input))
	}
}