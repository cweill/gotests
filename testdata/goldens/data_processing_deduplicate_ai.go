package testdata

import (
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test_deduplicate_empty_list",
			args: args{
				items: nil,
			},
			want: nil,
		},
		{
			name: "test_deduplicate_single_item",
			args: args{
				items: nil,
			},
			want: nil,
		},
		{
			name: "test_deduplicate_multiple_items",
			args: args{
				items: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		if got := Deduplicate(tt.args.items); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Deduplicate() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
