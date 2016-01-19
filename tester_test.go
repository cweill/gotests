package main

import (
	"tester/code"
	"tester/render"
	"testing"
)

func TestRenderRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			in: `testfiles/test1.go`,
			want: `func TestFoo1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Logf("Running: %v", tt.name)
		got := Foo1()
		if got != tt.want {
			t.Errorf("%v. Foo1() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test2.go`,
			want: `func TestFoo2(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Logf("Running: %v", tt.name)
		got := Foo2(tt.)
		if got != tt.want {
			t.Errorf("%v. Foo2() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		},
	}
	for _, tt := range tests {
		w := &fakeWriter{}
		render.TestCases(w, code.Parse(tt.in))
		got := string(w.got)
		if got != tt.want {
			t.Errorf("%v. TestCases() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

type fakeWriter struct {
	got []byte
}

func (f *fakeWriter) Write(p []byte) (n int, err error) {
	f.got = append(f.got, p...)
	return len(p), nil
}
