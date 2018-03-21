package options

type Options struct {
	PrintInputs    bool
	Subtests       bool
	ArgsStructMode ArgsStruct
}

type ArgsStruct string

const (
	ArgsStructAlways ArgsStruct = "always"
	ArgsStructNever  ArgsStruct = "never"
	ArgsStructSmart  ArgsStruct = "smart"
)
