package repeat

import "testing"

const WORD = "Hello"
const EMPTY_WORD = ""
const NEWLINE_WORD = "/n"

func TestRepeatTwice(t *testing.T) {
	repeated := repeat(WORD, 2)
	expected := "HelloHello"

	if repeated != expected {
		t.Errorf("Ожидаемое значение: %q. Получено: %q", expected, repeated)
	}
}

func TestRepeatTen(t *testing.T) {
	repeated := repeat(WORD, 10)
	expected := "HelloHelloHelloHelloHelloHelloHelloHelloHelloHello"

	if repeated != expected {
		t.Errorf("Ожидаемое значение: %q. Получено: %q", expected, repeated)
	}
}

func TestRepeatEmpty(t *testing.T) {
	repeated := repeat(EMPTY_WORD, 10)
	expected := ""

	if repeated != expected {
		t.Errorf("Ожидаемое значение: %q. Получено: %q", expected, repeated)
	}
}

func TestRepeatNewLine(t *testing.T) {
	repeated := repeat(NEWLINE_WORD, 4)
	expected := ""

	if repeated != expected {
		t.Errorf("Ожидаемое значение: %q. Получено: %q", expected, repeated)
	}
}