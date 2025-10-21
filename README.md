# gotests [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/cweill/gotests/blob/master/LICENSE) [![godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/cweill/gotests) [![Build Status](https://github.com/cweill/gotests/workflows/Go/badge.svg)](https://github.com/cweill/gotests/actions) [![Coverage Status](https://coveralls.io/repos/github/cweill/gotests/badge.svg?branch=master)](https://coveralls.io/github/cweill/gotests?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/cweill/gotests)](https://goreportcard.com/report/github.com/cweill/gotests)

`gotests` makes writing Go tests easy. It's a Golang commandline tool that generates [table driven tests](https://github.com/golang/go/wiki/TableDrivenTests) based on its target source files' function and method signatures. Any new dependencies in the test files are automatically imported.

## Demo

The following shows `gotests` in action using the [official Sublime Text 3 plugin](https://github.com/cweill/GoTests-Sublime). Plugins also exist for [Emacs](https://github.com/damienlevin/GoTests-Emacs), also [Emacs](https://github.com/s-kostyaev/go-gen-test), [Vim](https://github.com/buoto/gotests-vim), [Atom Editor](https://atom.io/packages/gotests), [Visual Studio Code](https://github.com/golang/vscode-go/blob/master/docs/settings.md#gogeneratetestsflags), and [IntelliJ Goland](https://www.jetbrains.com/help/go/run-debug-configuration-for-go-test.html).

![demo](https://github.com/cweill/GoTests-Sublime/blob/master/gotests.gif)

## Installation

__Minimum Go version:__ Go 1.22

Use [`go install`](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies) to install and update:
```sh
$ go install github.com/cweill/gotests/gotests@latest
```

## Usage

From the commandline, `gotests` can generate Go tests for specific source files or an entire directory. By default, it prints its output to `stdout`.

```sh
$ gotests [options] PATH ...
```

Available options:

```
  -all                  generate tests for all functions and methods

  -excl                 regexp. generate tests for functions and methods that don't
                         match. Takes precedence over -only, -exported, and -all

  -exported             generate tests for exported functions and methods. Takes
                         precedence over -only and -all

  -i                    print test inputs in error messages

  -named                switch table tests from using slice to map (with test name for the key)

  -only                 regexp. generate tests for functions and methods that match only.
                         Takes precedence over -all

  -nosubtests           disable subtest generation when >= Go 1.7

  -parallel             enable parallel subtest generation when >= Go 1.7.

  -w                    write output to (test) files instead of stdout

  -template_dir         Path to a directory containing custom test code templates. Takes
                         precedence over -template. This can also be set via environment
                         variable GOTESTS_TEMPLATE_DIR

  -template             Specify custom test code templates, e.g. testify. This can also
                         be set via environment variable GOTESTS_TEMPLATE

  -template_params_file read external parameters to template by json with file

  -template_params      read external parameters to template by json with stdin

  -use_go_cmp           use cmp.Equal (google/go-cmp) instead of reflect.DeepEqual

  -version              print version information and exit
```

## Contributions

Contributing guidelines are in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

`gotests` is released under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0).
