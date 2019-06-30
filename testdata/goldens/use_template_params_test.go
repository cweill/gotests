package use_template_params

import "testing"

func TestDo(t *testing.T) {
	// this is an example of how to use external parameters and custom templates
	type args struct {
		a    int
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		Do(tt.args.a, tt.args.name)
	}
}
