package testdata

import "testing"

func TestFoo5(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := Foo5()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Foo5() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Foo5() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkFoo5(b *testing.B) {
	benchmarks := []struct {
		name    string
		want    string
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		got, err := Foo5()
		if (err != nil) != bb.wantErr {
			b.Errorf("%q. Foo5() error = %v, wantErr %v", tt.name, err, bb.wantErr)
			continue
		}
		if got != bb.want {
			b.Errorf("%q. Foo5() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
