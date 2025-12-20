package week1

import (
	"fmt"
)

func helloWorld() string {
	return "hello, world"
}

func helloItsMe(name string) string {
	return fmt.Sprintf("hello, %s", name)
}
