# gotests [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests)
A Golang tool to automatically generate scaffolding for [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests).

The goal is to:
- [x] generate a framework for Go test functions in table driven style
- [x] automatically import test dependencies from file-under-test's
- [x] compare results using `==` on basic types and `reflect.DeepEqual` on everything else
- [ ] create fakes that conform to interfaces used in function parameters
- [ ] write test cases for you (_bluesky_)

## Example
Given the source file:
```Go
// testfiles/calculator.go
package testfiles

import "errors"

type Calculator struct{}

func (c *Calculator) Multiply(n, d int) int {
	return n * d
}

func (c *Calculator) Divide(n, d int) (int, error) {
	if d == 0 {
		return 0, errors.New("division by zero")
	}
	return n / d, nil
}

```
Running: 
```sh
$ gotests -w -i -all testfiles/calculator.go
Generated TestCalculatorMultiply
Generated TestCalculatorDivide
```
Generates the following test code:
```Go
// testfiles/calculator_test.go
package testfiles

import "testing"

func TestCalculatorMultiply(t *testing.T) {
	tests := []struct {
		n    int
		d    int
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &Calculator{}
		if got := c.Multiply(tt.n, tt.d); got != tt.want {
			t.Errorf("Calculator.Multiply(%v, %v) = %v, want %v", tt.n, tt.d, got, tt.want)
		}
	}
}

func TestCalculatorDivide(t *testing.T) {
	tests := []struct {
		n       int
		d       int
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &Calculator{}
		got, err := c.Divide(tt.n, tt.d)
		if (err != nil) != tt.wantErr {
			t.Errorf("Calculator.Divide(%v, %v) error = %v, wantErr %v", tt.n, tt.d, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("Calculator.Divide(%v, %v) = %v, want %v", tt.n, tt.d, got, tt.want)
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
To update from anywhere:
```sh
$ cwd=$(pwd); cd $GOPATH/src/github.com/cweill/gotests; git pull; go install; cd $cwd
```
## Usage
gotests only generates missing test functions, leaving existing ones intact. 
To generate only select tests for specific files, and output the results to stdout:
```sh
$ gotests -only="Foo|fetchBaz" foo.go bar.go
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
$ gotests -w -only="Foo|fetchBaz" foo.go bar.go # outputs to foo_test.go and bar_test.go
```
Now get that coverage up! 

## License

gotests is released under the [MIT License](http://www.opensource.org/licenses/MIT).
