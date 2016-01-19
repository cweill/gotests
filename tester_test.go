package main

import "testing"

func TestGenerateTestCases(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in: `testfiles/test1.go`,
			want: `package test1

import (
	"testing"
)

func TestFoo1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo1()
	}
}

`,
		}, {
			in: `testfiles/test2.go`,
			want: `package test2

import (
	"testing"
)

func TestFoo2(t *testing.T) {
	tests := []struct {
		name string
		in0 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo2(tt.in0)
	}
}

`,
		}, {
			in: `testfiles/test3.go`,
			want: `package test3

import (
	"testing"
)

func TestFoo3(t *testing.T) {
	tests := []struct {
		name string
		s string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Foo3(tt.s)
	}
}

`,
		}, {
			in: `testfiles/test4.go`,
			want: `package test4

import (
	"testing"
)

func TestFoo4(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got := Foo4()
		if got != tt.want {
			t.Errorf("%v. Foo4() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test5.go`,
			want: `package test5

import (
	"testing"
)

func TestFoo5(t *testing.T) {
	tests := []struct {
		name string
		want string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo5()
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo5() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%v. Foo5() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test6.go`,
			want: `package test6

import (
	"testing"
)

func TestFoo6(t *testing.T) {
	tests := []struct {
		name string
		i int
		b bool
		want string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo6(tt.i, tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo6() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%v. Foo6() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test7.go`,
			want: `package test7

import (
	"testing"
)

func TestFoo7(t *testing.T) {
	tests := []struct {
		name string
		b *Bar
		want string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.b.Foo7()
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo7() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%v. Foo7() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test8.go`,
			want: `package test8

import (
	"testing"
)

func TestFoo8(t *testing.T) {
	tests := []struct {
		name string
		b *Bar
		want *Bar
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo8(tt.b)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Foo8() error = %v, wantErr: %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo8() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test9.go`,
			want: `package test9

import (
	"testing"
)

func TestFoo9(t *testing.T) {
	tests := []struct {
		name string
		b Bar
		want Bar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got := tt.b.Foo9()
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v. Foo9() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		},
	}
	for _, tt := range tests {
		w := &logWriter{}
		generateTestCases(w, tt.in)
		got := string(w.log)
		if got != tt.want {
			t.Errorf("TestCases(%v) = %v, want %v", tt.in, got, tt.want)
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
