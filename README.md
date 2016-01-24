# gotests [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests)
A Go tool to automatically generate Go test code boilerplate in [table driven style](https://github.com/golang/go/wiki/TableDrivenTests).

The goal is to:
* generate missing Go test boilerplate for any specified functions and methods
* automatically import test dependencies from file-under-test's
* create fakes that conform to interfaces used in function parameters
* (_bluesky_) write test cases for you

## Example
Given the source file:
```Go
// testfiles/test007.go
package testfiles

type Bar struct{}

func (b *Bar) Foo7() (string, error) { return "", nil }
```
Running: 
```sh
$ gotests -w -only=Foo7 testfiles/test007.go
```
Generates the following test code:
```Go
// testfiles/test007_test.go
package testfiles

import "testing"

func TestFoo7(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		want    string
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
```
If the test file already exists, gotests generates and appends any non-existing tests. Any new dependencies are imported automatically.

## Installation
If your $GOPATH is setup just run:
```sh
$ go get github.com/cweill/gotests
```
Otherwise, setting up your $GOPATH is simple:
```sh
# consider adding it to your .bashrc or .bash_profile
$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
```
## Usage
gotests appends to existing test files or creates new ones next to the Go source files.

Generating only select tests for specific files and outputting the results to stdout:
```sh
$ gotests -only=Foo,fetchBaz foo.go bar.go
```
Or all tests:
```sh
$ gotests -all foo.go bar.go
```
Or most tests, excluding a few:
```sh
$ gotests -excl=fetchBaz foo.go bar.go
```
Generating tests for an entire directory:
```sh
$ gotests -all .
```
Passing the -w flag writes the output to the test files.
```sh
$ gotests -w -only=Foo,fetchBaz foo.go bar.go # outputs new tests to foo_test.go and bar_test.go
```
Now get that coverage up! 

## License

gotests is released under the [MIT License](http://www.opensource.org/licenses/MIT).
