package testdata

import "testing"

func TestBook_Open(t *testing.T) {
	tests := []struct {
		name    string
		b       *Book
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Book{}
		if err := b.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. Book.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func BenchmarkBook_Open(b *testing.B) {
	benchmarks := []struct {
		name    string
		b       *Book
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		b := &Book{}
		if err := b.Open(); (err != nil) != bb.wantErr {
			b.Errorf("%q. Book.Open() error = %v, wantErr %v", tt.name, err, bb.wantErr)
		}
	}
}

func Test_door_Open(t *testing.T) {
	tests := []struct {
		name    string
		d       *door
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		d := &door{}
		if err := d.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. door.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Benchmark_door_Open(b *testing.B) {
	benchmarks := []struct {
		name    string
		d       *door
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		d := &door{}
		if err := d.Open(); (err != nil) != bb.wantErr {
			b.Errorf("%q. door.Open() error = %v, wantErr %v", tt.name, err, bb.wantErr)
		}
	}
}

func Test_xml_Open(t *testing.T) {
	tests := []struct {
		name    string
		x       *xml
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		x := &xml{}
		if err := x.Open(); (err != nil) != tt.wantErr {
			t.Errorf("%q. xml.Open() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func Benchmark_xml_Open(b *testing.B) {
	benchmarks := []struct {
		name    string
		x       *xml
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		x := &xml{}
		if err := x.Open(); (err != nil) != bb.wantErr {
			b.Errorf("%q. xml.Open() error = %v, wantErr %v", tt.name, err, bb.wantErr)
		}
	}
}
