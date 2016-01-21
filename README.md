# gotests [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests)
A Go tool to automatically generate Go test code boilerplate.

The goal is to:
* generate missing Go test boilerplate for any specified functions and methods
* automatically import test dependencies from file-under-test's
* (optionally) generate fakes that conform to interfaces used in parameters
* (_bluesky_) generate test cases for you

## Example
Given the source file:
```Go
// testfiles/test007.go
package testfiles

type Bar struct{}

func (b *Bar) Foo7() (string, error) { return "", nil }
```
Running: 
```
$ gotests -funcs=Foo7 testfiles/test007.go
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
## Usage
Installation:
```
$ go get github.com/cweill/gotests
```
Generating only certain tests for specific files:
```
$ gotests -funcs=Foo,fetchBaz foo.go bar.go
```
Or all tests:
```
$ gotests -all foo.go bar.go
```
Or generating tests for an entire directory:
```
$ gotests -all .
```
Now get that coverage up! 

## License

gotests is released under the [MIT License](http://www.opensource.org/licenses/MIT).
