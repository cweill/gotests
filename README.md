# gotests [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/cweill/gotests/blob/master/LICENSE) [![Build Status](https://travis-ci.org/cweill/gotests.svg?branch=master)](https://travis-ci.org/cweill/gotests) [![Coverage Status](https://coveralls.io/repos/github/cweill/gotests/badge.svg?branch=master)](https://coveralls.io/github/cweill/gotests?branch=master) [![codebeat badge](https://codebeat.co/badges/7ef052e3-35ff-4cab-88f9-e13393c8ab35)](https://codebeat.co/projects/github-com-cweill-gotests)

`gotests` makes writing Go tests easy. It's a Golang commandline tool that generates [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) based on its target source files' function and method signatures. Additionally, any new dependencies in the test files are automatically imported.

> Writing good tests is not trivial, but in many situations a lot of ground can be covered with table-driven tests: Each table entry is a complete test case with inputs and expected results, and sometimes with additional information such as a test name to make the test output easily readable. If you ever find yourself using copy and paste when writing a test, think about whether refactoring into a table-driven test or pulling the copied code out into a helper function might be a better option. https://github.com/golang/go/wiki/TableDrivenTests#introduction

## Demo

The following shows `gotests` in action using the [official Sublime Text 3 plugin](https://github.com/cweill/GoTests-Sublime).

![demo](https://github.com/cweill/GoTests-Sublime/blob/master/gotests.gif)

There's a [plugin for Emacs](https://github.com/damienlevin/GoTests-Emacs) too.

## Installation

Use [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to install and update:

```sh
$ go get -u github.com/cweill/gotests/...
```

## Usage

From the commandline, `gotests` can generate tests for specific Go source files or an entire directory. By default, it prints its output to `stdout`.

```sh
$ gotests [options] PATH ...
```

Available options:

```
  -all         generate tests for all functions and methods
  
  -excl        regexp. generate tests for functions and methods that don't 
               match. Takes precedence over -only, -exported, and -all
    	   
  -exported    generate tests for exported functions and methods. Takes 
               precedence over -only and -all

  -i	       print test inputs in error messages
  
  -only        regexp. generate tests for functions and methods that match only.
               Takes precedence over -all
  
  -w           write output to (test) files instead of stdout
```

## License

`gotests` is released under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0).
