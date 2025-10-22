# gotests [![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/cweill/gotests/blob/master/LICENSE) [![godoc](https://img.shields.io/badge/go-documentation-blue.svg)](https://godoc.org/github.com/cweill/gotests) [![Build Status](https://github.com/cweill/gotests/workflows/Go/badge.svg)](https://github.com/cweill/gotests/actions) [![Coverage Status](https://coveralls.io/repos/github/cweill/gotests/badge.svg?branch=master)](https://coveralls.io/github/cweill/gotests?branch=master) [![codecov](https://codecov.io/gh/cweill/gotests/branch/master/graph/badge.svg)](https://codecov.io/gh/cweill/gotests) [![Go Report Card](https://goreportcard.com/badge/github.com/cweill/gotests)](https://goreportcard.com/report/github.com/cweill/gotests)

**gotests** is a powerful Go test generator that automatically creates [table-driven tests](https://go.dev/wiki/TableDrivenTests) from your source code. It analyzes function and method signatures to generate comprehensive test scaffolding, saving you time and ensuring consistency across your test suite.

### Key Features

- 🚀 **Zero-config test generation** - Works out of the box with any Go project
- 🎯 **Smart test scaffolding** - Generates complete table-driven test structure with proper type handling
- 🔧 **Flexible filtering** - Generate tests for specific functions, exported functions, or entire packages
- 📦 **Auto-imports** - Automatically adds required imports to test files
- 🧬 **Full generics support** - Works seamlessly with Go 1.18+ type parameters
- 🔄 **Recursive generation** - Process entire directory trees with `./...` pattern
- 🎨 **Custom templates** - Built-in support for testify and custom test templates
- ⚡ **Parallel subtests** - Optional parallel test execution support

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

  -ai                   generate test cases using AI (requires Ollama)

  -ai-model             AI model to use (default "qwen2.5-coder:0.5b")

  -ai-endpoint          Ollama API endpoint (default "http://localhost:11434")

  -ai-cases             number of test cases to generate with AI (default 3)

  -version              print version information and exit
```

## AI-Powered Test Generation

**gotests** can generate intelligent test cases using local LLMs via [Ollama](https://ollama.ai). This feature analyzes your function implementations and generates realistic test values, edge cases, and error conditions.

### Quick Start

1. **Install Ollama** ([https://ollama.ai](https://ollama.ai))

2. **Pull a model:**
   ```sh
   ollama pull qwen2.5-coder:0.5b  # Small, fast model (400MB)
   # or
   ollama pull llama3.2:latest     # Larger, more capable (2GB)
   ```

3. **Generate tests with AI:**
   ```sh
   gotests -all -ai -w yourfile.go
   ```

### Example

Given this function:
```go
func CalculateDiscount(price float64, percentage int) (float64, error) {
    if price < 0 {
        return 0, errors.New("price cannot be negative")
    }
    if percentage < 0 || percentage > 100 {
        return 0, errors.New("percentage must be between 0 and 100")
    }
    discount := price * float64(percentage) / 100.0
    return price - discount, nil
}
```

The AI generates:
```go
tests := []struct {
    name       string
    args       args
    want       float64
    wantErr    bool
}{
    {
        name: "valid_discount",
        args: args{price: 100.0, percentage: 20},
        want: 80.0,
        wantErr: false,
    },
    {
        name: "negative_price",
        args: args{price: -10.0, percentage: 20},
        want: 0,
        wantErr: true,
    },
    {
        name: "invalid_percentage",
        args: args{price: 100.0, percentage: 150},
        want: 0,
        wantErr: true,
    },
}
```

### AI Options

```sh
# Use a different model
gotests -all -ai -ai-model llama3.2:latest -w yourfile.go

# Generate more test cases
gotests -all -ai -ai-cases 5 -w yourfile.go

# Combine with other flags
gotests -exported -ai -parallel -w yourfile.go
```

### How It Works

- Analyzes function implementation and logic
- Generates realistic test values based on actual code
- Creates test cases for edge cases and error conditions
- Falls back to TODO comments if generation fails
- Works offline with local models (privacy-first)

### Supported Features

✅ Simple types (int, string, bool, float)
✅ Complex types (slices, maps, structs, pointers)
✅ Error returns and validation
✅ Variadic parameters
✅ Methods with receivers
✅ Multiple return values

### Privacy & Security

**What data is sent to the LLM:**
- Function signatures (name, parameters, return types)
- Complete function bodies including all code and comments
- No file paths or project context

**Privacy considerations:**
- ⚠️ **Function bodies may contain sensitive information** - business logic, algorithms, or credentials/secrets in comments
- ✅ **Local-first by default** - Using Ollama keeps all data on your machine; nothing is sent to external servers
- ✅ **Offline operation** - AI generation works completely offline with local models
- 🔒 **Recommendation**: Avoid using `-ai` on code containing secrets, API keys, or proprietary algorithms in comments

**If using cloud providers in the future:**
- Function source code will be transmitted to the cloud provider's API
- Review the provider's data retention and privacy policies
- Consider using `-ai` only on non-sensitive codebases

## Quick Start Examples

### Generate tests for a single function

Given a file `math.go`:

```go
package math

func Add(a, b int) int {
    return a + b
}
```

Generate a test:

```sh
$ gotests -only Add -w math.go
```

This creates `math_test.go` with:

```go
func TestAdd(t *testing.T) {
    type args struct {
        a int
        b int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.args.a, tt.args.b); got != tt.want {
                t.Errorf("Add() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Generate tests for all exported functions

```sh
$ gotests -all -exported -w .
```

This generates tests for all exported functions in the current directory.

### Generate tests recursively

```sh
$ gotests -all -w ./...
```

This generates tests for all functions in the current directory and all subdirectories.

### Generate tests with testify

```sh
$ gotests -all -template testify -w calculator.go
```

This generates tests using the [testify](https://github.com/stretchr/testify) assertion library.

### Common workflows

```sh
# Generate tests for all exported functions in a package
$ gotests -exported -w pkg/*.go

# Generate only tests for specific functions matching a pattern
$ gotests -only "^Process" -w handler.go

# Generate tests excluding certain functions
$ gotests -all -excl "^helper" -w utils.go

# Generate parallel subtests
$ gotests -all -parallel -w service.go
```

## Go Generics Support

`gotests` fully supports Go generics (type parameters) introduced in Go 1.18+. It automatically generates tests for generic functions and methods on generic types.

### Example: Generic Function

Given this generic function:

```go
func FindFirst[T comparable](slice []T, target T) (int, error) {
    for i, v := range slice {
        if v == target {
            return i, nil
        }
    }
    return -1, ErrNotFound
}
```

Running `gotests -all -w yourfile.go` generates:

```go
func TestFindFirst(t *testing.T) {
    type args struct {
        slice  []string
        target string
    }
    tests := []struct {
        name    string
        args    args
        want    int
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := FindFirst[string](tt.args.slice, tt.args.target)
            if (err != nil) != tt.wantErr {
                t.Errorf("FindFirst() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("FindFirst() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Example: Methods on Generic Types

`gotests` also supports methods on generic types:

```go
type Set[T comparable] struct {
    m map[T]struct{}
}

func (s *Set[T]) Add(v T) {
    if s.m == nil {
        s.m = make(map[T]struct{})
    }
    s.m[v] = struct{}{}
}
```

Generates:

```go
func TestSet_Add(t *testing.T) {
    type args struct {
        v string
    }
    tests := []struct {
        name string
        s    *Set[string]
        args args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            s := &Set[string]{}
            s.Add(tt.args.v)
        })
    }
}
```

### Type Constraint Mapping

`gotests` uses intelligent defaults for type parameter instantiation:

- `any` → `int`
- `comparable` → `string`
- Union types (`int64 | float64`) → first option (`int64`)
- Approximation constraints (`~int`) → underlying type (`int`)

This ensures generated tests use appropriate concrete types for testing generic code.

## Contributions

Contributing guidelines are in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

`gotests` is released under the [Apache 2.0 License](http://www.apache.org/licenses/LICENSE-2.0).
