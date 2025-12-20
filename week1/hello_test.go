package week1

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	msg := helloWorld()
	if msg != "hello, world" {
		t.Errorf("expected 'hello, world' but received '%s'", msg)
	}
}

func TestHelloItsMe(t *testing.T) {
	names := []string{"tom", "mara", "bix"}
	for _, name := range names {
		msg := helloItsMe(name)
		if msg != fmt.Sprintf("hello, %s", name) {
			t.Errorf("expected 'hello, %s' but received '%s'", name, msg)
		}
	}
}
