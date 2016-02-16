# gotests [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/cweill/gotests/blob/master/LICENSE) [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests) [![Coverage Status](https://coveralls.io/repos/github/cweill/gotests/badge.svg?branch=master)](https://coveralls.io/github/cweill/gotests?branch=master)

`gotests` is a Golang commandline tool for automatically generating [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests). For a given source file, `gotests` generates missing tests based on its function and method signatures. Any new dependencies in the test file are automatically imported.

## Demo

The following demo shows Sublime Text 3 integration with `gotests`.

![demo](/editors/SublimeText3/gotests.gif)

Plugins for emacs and vim are coming soon.

## Installation

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:
```sh
$ go get -u github.com/cweill/gotests/...
```

## Usage

From the commandline, `gotests` can generate tests for specific Go source files or an entire directory. By default, it prints its output to stdout.
```sh
$ gotests [options] PATH ...
```
Available options:
```
  -all     generate tests for all functions and methods
  
  -excl    regexp. generate tests for functions and methods that don't match. e.g. -excl="^\p{Ll}" filters unexported
    	   functions and methods only. Takes precedence over -only and -all
    	   
  -i	   print test inputs in error messages
  
  -only    regexp. generate tests for functions and methods that match only. e.g. -only="^\p{Lu}" selects exported 
           functions and methods only. Takes precedence over -all
  
  -w       write output to (test) files instead of stdout
```

## License

gotests is released under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0).
