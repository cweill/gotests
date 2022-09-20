package main

func Foo(input bool) bool {
	return input
}

func Bar() bool {
	return Foo(true)
}