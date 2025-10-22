// A commandline tool for generating table-driven Go tests.
//
// This tool can generate tests for specific Go source files or an entire
// directory. By default, it prints its output to stdout.
//
// Usage:
//
//   $ gotests [options] PATH ...
//
// Available options:
//
//   -all                  generate tests for all functions and methods
//
//   -excl                 regexp. generate tests for functions and methods that don't
//                         match. Takes precedence over -only, -exported, and -all
//
//   -exported             generate tests for exported functions and methods. Takes
//                         precedence over -only and -all
//
//   -i                    print test inputs in error messages
//
//   -named                switch table tests from using slice to map (with test name for the key)
//
//   -only                 regexp. generate tests for functions and methods that match only.
//                         Takes precedence over -all
//
//   -nosubtests           disable generating tests using the Go 1.7 subtests feature
//
//   -parallel             enable generating parallel subtests using the Go 1.7 feature
//
//   -w                    write output to (test) files instead of stdout
//
//   -template_dir         Path to a directory containing custom test code templates. Takes
//                         precedence over -template. This can also be set via environment
//                         variable GOTESTS_TEMPLATE_DIR
//
//   -template             Specify custom test code templates, e.g. testify. This can also
//                         be set via environment variable GOTESTS_TEMPLATE
//
//   -template_params_file read external parameters to template by json with file
//
//   -template_params      read external parameters to template by json with stdin
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/cweill/gotests/gotests/process"
)

var (
	onlyFuncs          = flag.String("only", "", `regexp. generate tests for functions and methods that match only. Takes precedence over -all`)
	exclFuncs          = flag.String("excl", "", `regexp. generate tests for functions and methods that don't match. Takes precedence over -only, -exported, and -all`)
	exportedFuncs      = flag.Bool("exported", false, `generate tests for exported functions and methods. Takes precedence over -only and -all`)
	allFuncs           = flag.Bool("all", false, "generate tests for all functions and methods")
	printInputs        = flag.Bool("i", false, "print test inputs in error messages")
	writeOutput        = flag.Bool("w", false, "write output to (test) files instead of stdout")
	templateDir        = flag.String("template_dir", "", `optional. Path to a directory containing custom test code templates. Takes precedence over -template. This can also be set via environment variable GOTESTS_TEMPLATE_DIR`)
	template           = flag.String("template", "", `optional. Specify custom test code templates, e.g. testify. This can also be set via environment variable GOTESTS_TEMPLATE`)
	templateParamsPath = flag.String("template_params_file", "", "read external parameters to template by json with file")
	templateParams     = flag.String("template_params", "", "read external parameters to template by json with stdin")
	useGoCmp           = flag.Bool("use_go_cmp", false, "use cmp.Equal (google/go-cmp) instead of reflect.DeepEqual")
	useAI              = flag.Bool("ai", false, "generate test cases using AI (requires Ollama)")
	aiModel            = flag.String("ai-model", "qwen2.5-coder:0.5b", "AI model to use for test generation")
	aiEndpoint         = flag.String("ai-endpoint", "http://localhost:11434", "Ollama API endpoint")
	aiCases            = flag.Int("ai-cases", 3, "number of test cases to generate with AI")
	version            = flag.Bool("version", false, "print version information and exit")
)

var (
	// nosubtests is always set to default value of true when Go < 1.7.
	// When >= Go 1.7 the default value is changed to false by the
	// flag.BoolVar but can be overridden by setting nosubtests to true
	nosubtests = true

	// parallel is default false.
	parallel bool

	// use map instead of slice for table tests.
	named bool
)

func main() {
	flag.Parse()
	args := flag.Args()

	if *version {
		printVersion()
		return
	}

	// Validate AI parameters and warn user
	if *useAI {
		// Warn about sending code to AI provider
		fmt.Fprintf(os.Stderr, "⚠️  WARNING: Function source code will be sent to AI provider at %s\n", *aiEndpoint)
		fmt.Fprintf(os.Stderr, "   Ensure your code does not contain secrets or sensitive information.\n\n")

		// Validate parameters
		if *aiModel == "" {
			fmt.Fprintf(os.Stderr, "Error: -ai-model cannot be empty when using -ai flag\n")
			os.Exit(1)
		}
		if *aiCases < 1 || *aiCases > 100 {
			fmt.Fprintf(os.Stderr, "Error: -ai-cases must be between 1 and 100, got %d\n", *aiCases)
			os.Exit(1)
		}
	}

	process.Run(os.Stdout, args, &process.Options{
		OnlyFuncs:          *onlyFuncs,
		ExclFuncs:          *exclFuncs,
		ExportedFuncs:      *exportedFuncs,
		AllFuncs:           *allFuncs,
		PrintInputs:        *printInputs,
		Subtests:           !nosubtests,
		Parallel:           parallel,
		Named:              named,
		WriteOutput:        *writeOutput,
		Template:           valOrGetenv(*template, "GOTESTS_TEMPLATE"),
		TemplateDir:        valOrGetenv(*templateDir, "GOTESTS_TEMPLATE_DIR"),
		TemplateParamsPath: *templateParamsPath,
		TemplateParams:     *templateParams,
		UseGoCmp:           *useGoCmp,
		UseAI:              *useAI,
		AIModel:            *aiModel,
		AIEndpoint:         *aiEndpoint,
		AICases:            *aiCases,
	})
}

func printVersion() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("gotests (unknown version)")
		return
	}

	version := info.Main.Version
	if version == "" || version == "(devel)" {
		version = "development"
	}

	fmt.Printf("gotests %s\n", version)
	fmt.Printf("Go version: %s\n", info.GoVersion)

	// Print VCS information if available
	for _, setting := range info.Settings {
		if setting.Key == "vcs.revision" {
			fmt.Printf("Git commit: %s\n", setting.Value)
		}
		if setting.Key == "vcs.time" {
			fmt.Printf("Build time: %s\n", setting.Value)
		}
	}
}

func valOrGetenv(val, key string) string {
	if val != "" {
		return val
	}
	return os.Getenv(key)
}
