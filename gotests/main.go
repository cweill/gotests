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
//   -only                 regexp. generate tests for functions and methods that match only.
//                         Takes precedence over -all
//
//   -nosubtests           disable subtest generation when >= Go 1.7
//
//   -parallel             enable parallel subtest generation when >= Go 1.7.
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
//   -template_params	   read external parameters to template by json with stdin
package main

import (
	"flag"
	"os"

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
)

var (
	// nosubtests is always set to default value of true when Go < 1.7.
	// When >= Go 1.7 the default value is changed to false by the
	// flag.BoolVar but can be overridden by setting nosubtests to true
	nosubtests = true

	// parallel is default false.
	parallel bool
)

func main() {
	flag.Parse()
	args := flag.Args()

	process.Run(os.Stdout, args, &process.Options{
		OnlyFuncs:          *onlyFuncs,
		ExclFuncs:          *exclFuncs,
		ExportedFuncs:      *exportedFuncs,
		AllFuncs:           *allFuncs,
		PrintInputs:        *printInputs,
		Subtests:           !nosubtests,
		Parallel:           parallel,
		WriteOutput:        *writeOutput,
		Template:           valOrGetenv(*template, "GOTESTS_TEMPLATE"),
		TemplateDir:        valOrGetenv(*templateDir, "GOTESTS_TEMPLATE_DIR"),
		TemplateParamsPath: *templateParamsPath,
	})
}

func valOrGetenv(val, key string) string {
	if val != "" {
		return val
	}
	return os.Getenv(key)
}
