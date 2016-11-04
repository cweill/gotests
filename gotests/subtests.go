// +build go1.7

package main

import "flag"

func init() {
	flag.BoolVar(&subtests, "nosubtests", false, "disable generating tests using the Go 1.7 subtests feature")
}
