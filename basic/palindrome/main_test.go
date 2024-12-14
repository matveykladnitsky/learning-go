package palindrome

import "testing"

func TestIsPalindromeTrue(t *testing.T) {
	word := "madam"
	if !isPalindrome(word) {
		t.Errorf("Ожидается true, но получено false для слова: %q", word)
	}
}

func TestIsPalindromeFalse(t *testing.T) {
	word := "hello"
	if isPalindrome(word) {
		t.Errorf("Ожидается false, но получено true для слова: %q", word)
	}
}

func TestIsPalindromeEmpty(t *testing.T) {
	word := ""
	if !isPalindrome(word) {
		t.Errorf("Ожидается true, но получено false для пустой строки")
	}
}

func TestIsPalindromeSingleChar(t *testing.T) {
	word := "a"
	if !isPalindrome(word) {
		t.Errorf("Ожидается true, но получено false для слова: %q", word)
	}
}