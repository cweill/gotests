package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/cweill/gotests"
)

const newFilePerm os.FileMode = 0644

var (
	onlyFuncs     = flag.String("only", "", `regexp. generate tests for functions and methods that match only. Takes precedence over -all`)
	exclFuncs     = flag.String("excl", "", `regexp. generate tests for functions and methods that don't match. Takes precedence over -only, -exp, and -all`)
	exportedFuncs = flag.Bool("exported", false, `generate tests for exported functions and methods. Takes precedence over -only and -all`)
	allFuncs      = flag.Bool("all", false, "generate tests for all functions and methods")
	printInputs   = flag.Bool("i", false, "print test inputs in error messages")
	writeOutput   = flag.Bool("w", false, "write output to (test) files instead of stdout")
)

func main() {
	flag.Parse()
	if *onlyFuncs == "" && *exclFuncs == "" && !*exportedFuncs && !*allFuncs {
		fmt.Println("Please specify either the -only, -excl, -exp, or -all flag")
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
	for _, path := range flag.Args() {
		tests, err := gotests.GenerateTests(path, &gotests.Options{
			Only:        onlyRE,
			Exclude:     exclRE,
			Exported:    *exportedFuncs,
			PrintInputs: *printInputs,
		})
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if len(tests) == 0 {
			fmt.Println("No tests generated for", path)
			continue
		}
		for _, test := range tests {
			if *writeOutput {
				if err := ioutil.WriteFile(test.Path, test.Output, newFilePerm); err != nil {
					fmt.Println(err)
					continue
				}
			}
			for _, t := range test.Functions {
				fmt.Printf("Generated %v\n", t.TestName())
			}
			if !*writeOutput {
				fmt.Println(string(test.Output))
			}
		}
	}
}
