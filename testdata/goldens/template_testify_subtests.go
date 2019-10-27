package testdata

import "testing"

func TestFoo201a(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Foo201a())
		})
	}
}

func TestFoo201b(t *testing.T) {
	tests := []struct {
		name      string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, Foo201b())
		})
	}
}

func TestFoo201c(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, Foo201c(tt.args.n))
		})
	}
}

func TestFoo201d(t *testing.T) {
	tests := []struct {
		name      string
		want      bool
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Foo201d()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFoo201e(t *testing.T) {
	type args struct {
		n int
		s string
	}
	tests := []struct {
		name      string
		args      args
		want      bool
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Foo201e(tt.args.n, tt.args.s)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
