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

  -ai                   generate test cases using AI

  -ai-provider          AI provider: ollama, openai, claude, gemini (default "ollama")

  -ai-model             AI model to use (default depends on provider)

  -ai-endpoint          AI API endpoint (default depends on provider)

  -ai-api-key           API key for cloud AI providers (openai, claude, gemini).
                        Can also be set via GOTESTS_AI_API_KEY environment variable

  -ai-min-cases         minimum number of test cases to generate with AI (default 3)

  -ai-max-cases         maximum number of test cases to generate with AI (default 10)

  -version              print version information and exit
```

## AI-Powered Test Generation

**gotests** can generate intelligent test cases using AI. This feature analyzes your function implementations and generates realistic test values, edge cases, and error conditions.

### Supported AI Providers

| Provider | Type | Models | Pricing |
|----------|------|--------|---------|
| **Ollama** | Local | qwen2.5-coder, llama3.2, codellama | Free (runs locally) |
| **OpenAI** | Cloud | gpt-4o, gpt-4o-mini, gpt-4-turbo | Pay-per-token |
| **Claude** | Cloud | claude-3-opus, claude-3-sonnet, claude-3-haiku | Pay-per-token |
| **Gemini** | Cloud | gemini-1.5-pro, gemini-1.5-flash | Pay-per-token |
| **Any OpenAI-compatible** | Cloud/Local | Various | Varies |

### OpenAI-Compatible API Support

The `openai` provider works with **any service** that implements the OpenAI Chat Completions API format. This includes:

| Service | Endpoint | Notes |
|---------|----------|-------|
| OpenAI | `https://api.openai.com/v1` | Default |
| Azure OpenAI | `https://{resource}.openai.azure.com/...` | Enterprise |
| Together AI | `https://api.together.xyz/v1` | Open-source models |
| Groq | `https://api.groq.com/openai/v1` | Ultra-fast inference |
| Anyscale | `https://api.endpoints.anyscale.com/v1` | Open-source models |
| Fireworks AI | `https://api.fireworks.ai/inference/v1` | Fast inference |
| DeepSeek | `https://api.deepseek.com/v1` | Coding models |
| Mistral AI | `https://api.mistral.ai/v1` | European provider |
| Perplexity | `https://api.perplexity.ai` | Search-augmented |
| OpenRouter | `https://openrouter.ai/api/v1` | Multi-provider gateway |
| LM Studio | `http://localhost:1234/v1` | Local GUI app |

### Quick Start with Ollama (Local, Free)

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

### Quick Start with Cloud Providers

#### OpenAI

```sh
# Set your API key
export GOTESTS_AI_API_KEY="sk-..."

# Generate tests with GPT-4o-mini (fast and cheap)
gotests -all -ai -ai-provider openai -w yourfile.go

# Use a specific model
gotests -all -ai -ai-provider openai -ai-model gpt-4o -w yourfile.go
```

#### Any OpenAI-Compatible Provider

```sh
# Together AI
export GOTESTS_AI_API_KEY="your-together-key"
gotests -all -ai -ai-provider openai -ai-endpoint https://api.together.xyz/v1 -ai-model meta-llama/Llama-3-70b-chat-hf -w yourfile.go

# Groq (ultra-fast)
export GOTESTS_AI_API_KEY="your-groq-key"
gotests -all -ai -ai-provider openai -ai-endpoint https://api.groq.com/openai/v1 -ai-model llama-3.1-70b-versatile -w yourfile.go

# DeepSeek (coding specialist)
export GOTESTS_AI_API_KEY="your-deepseek-key"
gotests -all -ai -ai-provider openai -ai-endpoint https://api.deepseek.com/v1 -ai-model deepseek-coder -w yourfile.go

# LM Studio (local, no API key needed)
gotests -all -ai -ai-provider openai -ai-endpoint http://localhost:1234/v1 -ai-model local-model -w yourfile.go

# OpenRouter (access multiple providers)
export GOTESTS_AI_API_KEY="your-openrouter-key"
gotests -all -ai -ai-provider openai -ai-endpoint https://openrouter.ai/api/v1 -ai-model anthropic/claude-3-sonnet -w yourfile.go
```

#### Claude (Anthropic Native API)

```sh
# Set your API key
export GOTESTS_AI_API_KEY="sk-ant-..."

# Generate tests with Claude
gotests -all -ai -ai-provider claude -w yourfile.go

# Use a specific model
gotests -all -ai -ai-provider claude -ai-model claude-3-sonnet-20240229 -w yourfile.go
```

#### Gemini (Google Native API)

```sh
# Set your API key
export GOTESTS_AI_API_KEY="AIza..."

# Generate tests with Gemini
gotests -all -ai -ai-provider gemini -w yourfile.go

# Use a specific model
gotests -all -ai -ai-provider gemini -ai-model gemini-1.5-pro -w yourfile.go
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

The AI generates (showing 3 cases; by default, the AI generates between 3-10 cases):
```go
func TestCalculateDiscount(t *testing.T) {
    type args struct {
        price      float64
        percentage int
    }
    tests := []struct {
        name    string
        args    args
        want    float64
        wantErr bool
    }{
        {
            name: "valid discount",
            args: args{price: 100.0, percentage: 20},
            want: 80.0,
            wantErr: false,
        },
        {
            name: "negative price",
            args: args{price: -10.0, percentage: 20},
            want: 0,
            wantErr: true,
        },
        {
            name: "invalid percentage",
            args: args{price: 100.0, percentage: 150},
            want: 0,
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := CalculateDiscount(tt.args.price, tt.args.percentage)
            if (err != nil) != tt.wantErr {
                t.Errorf("CalculateDiscount() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("CalculateDiscount() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### AI Options

```sh
# Use Ollama with a different model
gotests -all -ai -ai-model llama3.2:latest -w yourfile.go

# Use OpenAI with API key from environment
export GOTESTS_AI_API_KEY="sk-..."
gotests -all -ai -ai-provider openai -ai-model gpt-4o -w yourfile.go

# Use Claude with API key as flag
gotests -all -ai -ai-provider claude -ai-api-key "sk-ant-..." -w yourfile.go

# Use Gemini
gotests -all -ai -ai-provider gemini -ai-api-key "AIza..." -w yourfile.go

# Generate a specific number of test cases (min = max)
gotests -all -ai -ai-min-cases 5 -ai-max-cases 5 -w yourfile.go

# Generate a range of test cases (AI chooses between 3-7)
gotests -all -ai -ai-min-cases 3 -ai-max-cases 7 -w yourfile.go

# Combine with other flags
gotests -exported -ai -parallel -w yourfile.go
```

### Default Models by Provider

| Provider | Default Model | Default Endpoint |
|----------|---------------|------------------|
| Ollama | `qwen2.5-coder:0.5b` | `http://localhost:11434` |
| OpenAI | `gpt-4o-mini` | `https://api.openai.com/v1` |
| Claude | `claude-3-haiku-20240307` | `https://api.anthropic.com/v1` |
| Gemini | `gemini-1.5-flash` | `https://generativelanguage.googleapis.com/v1beta` |

**Note:** When using `-ai-provider openai` with a custom `-ai-endpoint`, you must also specify `-ai-model` as different providers support different models.

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
