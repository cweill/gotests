# gotests [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests)
A Go tool to automatically generate Go test code boilerplate in [table driven style](https://github.com/golang/go/wiki/TableDrivenTests).

The goal is to:
- [x] generate missing Go test boilerplate for any specified functions and methods
- [x] automatically import test dependencies from file-under-test's
- [x] compare results using `==` on basic types and `reflect.DeepEqual` on everything else
- [ ] create fakes that conform to interfaces used in function parameters
- [ ] write test cases for you (_bluesky_)

## Example
Given the source file:
```Go
// testfiles/test007.go
package testfiles

type Bar struct{}

func (b *Bar) Foo7(i int) (string, error) {
  return "", nil
}
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

func TestBarFoo7(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		i       int
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.b.Foo7(tt.i)
		if (err != nil) != tt.wantErr {
			t.Errorf("%v. Bar.Foo7() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%v. Bar.Foo7() = %v, want %v", tt.name, got, tt.want)
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
$ mkdir $HOME/go
# consider adding below to your .bashrc or .bash_profile
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
```
To update:
```sh
$ cd $GOPATH/src/github.com/cweill/gotests
$ git pull
$ go install
```
## Usage
gotests only generates missing test functions, leaving existing ones intact. 
To generate only select tests for specific files, and output the results to stdout:
```sh
$ gotests -only=Foo,fetchBaz foo.go bar.go
```
For all tests:
```sh
$ gotests -all foo.go bar.go
```
For most tests, excluding a few:
```sh
$ gotests -excl=fetchBaz foo.go bar.go
```
To generate tests for an entire directory:
```sh
$ gotests -all .
```
Pass the -w flag to write the output to the test files. gotests appends to existing test files or creates new ones beside the source files.
```sh
$ gotests -w -only=Foo,fetchBaz foo.go bar.go # outputs to foo_test.go and bar_test.go
```
Now get that coverage up! 

## License

gotests is released under the [MIT License](http://www.opensource.org/licenses/MIT).
