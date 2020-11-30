package main

import (
	"os"
	"testing"
)

func Test_valOrGetenv(t *testing.T) {
	const (
		testEnvVarKey = "GOTESTS_TEST_KEY"
		testEnvVarVal = "testify"
	)

	if err := os.Setenv(testEnvVarKey, testEnvVarVal); err != nil {
		t.Fatalf("setting environment variable: %v", err)
	}

	defer func(t *testing.T) {
		if err := os.Unsetenv(testEnvVarKey); err != nil {
			t.Fatalf("unsetting environment variable: %v", err)
		}
	}(t)

	type args struct {
		val string
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "both val and key are empty",
			args: args{
				val: "",
				key: "",
			},
			want: "",
		},
		{
			name: "val=my_template, key=GOTESTS_TEST_KEY",
			args: args{
				val: "my_template",
				key: testEnvVarKey,
			},
			want: "my_template",
		},
		{
			name: "val is empty, key=GOTESTS_TEST_KEY",
			args: args{
				val: "",
				key: testEnvVarKey,
			},
			want: testEnvVarVal,
		},
		{
			name: "val is empty, key contains unset key",
			args: args{
				val: "",
				key: "GOTESTS_TEST_KEY_X",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valOrGetenv(tt.args.val, tt.args.key); got != tt.want {
				t.Errorf("valOrGetenv() = %v, want %v", got, tt.want)
			}
		})
	}
}
