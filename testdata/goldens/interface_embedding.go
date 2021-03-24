package undefinedtypes

import "testing"

func TestSomeStruct_Do(t *testing.T) {
	type fields struct {
		Doer some.Doer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &SomeStruct{
			Doer: tt.fields.Doer,
		}
		c.Do()
	}
}
