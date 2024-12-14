package helloworld

import "testing"


func TestHelloWorld(t *testing.T) {
	repeated := sayHelloWorld()
	expected := "Hello World!"

	if repeated != expected {
		t.Errorf("Ожидаемое значение: %q. Получено: %q", expected, repeated)
	}
}