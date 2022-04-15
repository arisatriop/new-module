package gomodule2

import "fmt"

func SayHi(name string, me string) string {
	return fmt.Sprintf("Hi, %s my name is %s.\n", name, me)
}

func SayHiWithName(name string) string {
	return fmt.Sprintf("Hi, %s.\n", name)
}
