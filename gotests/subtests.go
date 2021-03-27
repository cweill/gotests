// +build go1.7

package main

import "flag"

func init() {
	flag.BoolVar(&nosubtests, "nosubtests", false, "disable generating tests using the Go 1.7 subtests feature")
	flag.BoolVar(&parallel, "parallel", false, "enable generating parallel subtests using the Go 1.7 feature")
	flag.BoolVar(&named, "named", false, "switch table tests from using slice to map (with test name for the key)")
}
