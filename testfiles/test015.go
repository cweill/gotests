package test15

func Foo15(f func(string) string) func(int) int {
	return func(int) int { return 1 }
}
