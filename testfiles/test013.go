package test13

func Foo13(f func()) func() { return func() {} }
