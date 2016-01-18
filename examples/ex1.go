package ex1

func Foo1() {}

func Foo2(string) {}

func Foo3(s string) {}

func Foo4() string { return "" }

func Foo5() (string, error) { return "", nil }

func Foo6() (string, error) { return "", nil }

type Bar struct{}

func (b *Bar) Foo7() (string, error) { return "", nil }

func (b *Bar) Foo8(i int) (string, error) { return "", nil }
