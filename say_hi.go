package gomodule2

import "fmt"

func SayHi() string {
	return "Hi..."
}

func SayHiWithName(name string) string {
	return fmt.Sprintf("Hi, %s.\n", name)
}
