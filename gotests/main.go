package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/cweill/gotests"
	"github.com/cweill/gotests/input"
)

var (
	onlyFuncs   = flag.String("only", "", `regexp. generate tests for functions and methods that match only. e.g. -only="^\p{Lu}" selects exported functions and methods only. Takes precedence over -all`)
	exclFuncs   = flag.String("excl", "", `regexp. generate tests for functions and methods that don't match. e.g. -excl="^\p{Ll}" filters unexported functions and methods only. Takes precedence over -only and -all`)
	allFuncs    = flag.Bool("all", false, "generate tests for all functions and methods")
	printInputs = flag.Bool("i", false, "print test inputs in error messages")
	writeOutput = flag.Bool("w", false, "write output to (test) files instead of stdout")
)

func main() {
	flag.Parse()
	if *onlyFuncs == "" && *exclFuncs == "" && !*allFuncs {
		fmt.Println("Please specify either the -only, -excl, or -all flag")
		return
	}
	if len(flag.Args()) == 0 {
		fmt.Println("Please specify a file or directory containing the source")
		return
	}
	var onlyRE, exclRE *regexp.Regexp
	var err error
	if *onlyFuncs != "" {
		onlyRE, err = regexp.Compile(*onlyFuncs)
		if err != nil {
			fmt.Printf("Invalid -only regex: %v\n", err)
			return
		}
	}
	if *exclFuncs != "" {
		exclRE, err = regexp.Compile(*exclFuncs)
		if err != nil {
			fmt.Printf("Invalid -excl regex: %v\n", err)
			return
		}
	}
	var count int
	for _, path := range flag.Args() {
		ps, err := input.Files(path)
		if err != nil {
			if err == input.NoFilesFound {
				fmt.Printf("No source files found at %v\n", path)
			} else {
				fmt.Println(err.Error())
			}
			continue
		}
		for _, src := range ps {
			tests, b, err := gotests.GenerateTests(string(src), src.TestPath(), src.TestPath(), &gotests.Options{
				Only:        onlyRE,
				Excl:        exclRE,
				Write:       *writeOutput,
				PrintInputs: *printInputs,
			})
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if len(tests) == 0 {
				continue
			}
			for _, test := range tests {
				fmt.Printf("Generated %v\n", test.TestName())
				count++
			}
			if !*writeOutput {
				fmt.Println(string(b))
			}
		}
	}
	if count == 0 {
		fmt.Println("No tests generated")
	}
}
