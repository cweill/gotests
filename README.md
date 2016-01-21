# gotests [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests)
A tool to generate test code boilerplate for exported Golang functions and methods.

## Usage
Installation:
```
$ go get github.com/cweill/gotests
```
Generating tests for specific files:
```
$ gotest  my/source/dir/foo.go  my/source/dir/bar.go
```
You can also generate tests for an entire directory:
```
$ gotest my/source/dir
```
Now get that coverage up! 

## License

gotests is released under the [MIT License](http://www.opensource.org/licenses/MIT).
