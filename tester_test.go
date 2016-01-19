package main

import "testing"

func TestGenerateTestCases(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in: `testfiles/test1.go`,
			want: `package testfiles

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
			want: `package testfiles

import (
	"testing"
)

func TestFoo2(t *testing.T) {
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
		}, {
			in: `testfiles/test3.go`,
			want: `package testfiles

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
		t.Logf("Running: %v", tt.name)
		got := Foo3(tt.s)
		if got != tt.want {
			t.Errorf("%v. Foo3() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test4.go`,
			want: `package testfiles

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
		t.Logf("Running: %v", tt.name)
		got := Foo4()
		if got != tt.want {
			t.Errorf("%v. Foo4() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

`,
		}, {
			in: `testfiles/test5.go`,
			want: `package testfiles

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
		t.Logf("Running: %v", tt.name)
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
			want: `package testfiles

import (
	"testing"
)

func TestFoo6(t *testing.T) {
	tests := []struct {
		name string
		i int
		want string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Logf("Running: %v", tt.name)
		got, err := Foo6(tt.i)
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
