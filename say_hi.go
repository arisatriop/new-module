package gomodule2

import "fmt"

func SayHi(name string) string {
	return fmt.Sprintf("Hi..., %s.\n", name)
}

func SayHiWithName(name string) string {
	return fmt.Sprintf("Hi, %s.\n", name)
}
