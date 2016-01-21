# gotests [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests)
A Go tool to generate test code boilerplate for exported Golang functions and methods.

The goal is to:
* generate missing Go test boilerplate for __exported__ functions and methods
* automatically import test dependencies from file-under-test's
* (optionally) generate fakes that conform to interfaces used in parameters
* (_bluesky_) generate test cases for you

## Example
Given the source file:
```Go
// testfiles/test007.go
package test7

type Bar struct{}

func (b *Bar) Foo7() (string, error) { return "", nil }
```
Running: 
```
$ gotests testfiles/test007.go
```
Generates the following test code:
```Go
// testfiles/test007_test.go
package test7

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
Generating tests for specific files:
```
$ gotests  my/source/dir/foo.go  my/source/dir/bar.go
```
You can also generate tests for an entire directory:
```
$ gotests my/source/dir
```
Now get that coverage up! 

## License

gotests is released under the [MIT License](http://www.opensource.org/licenses/MIT).
