package testdata

import (
	"bytes"
	"testing"
)

func TestBar_Write(t *testing.T) {
	tests := []struct {
		name    string
		b       *Bar
		wantW   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		b := &Bar{}
		w := &bytes.Buffer{}
		if err := b.Write(w); (err != nil) != tt.wantErr {
			t.Errorf("%q. Bar.Write() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if gotW := w.String(); gotW != tt.wantW {
			t.Errorf("%q. Bar.Write() = %v, want %v", tt.name, gotW, tt.wantW)
		}
	}
}

func BenchmarkBar_Write(b *testing.B) {
	benchmarks := []struct {
		name    string
		b       *Bar
		wantW   string
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		b := &Bar{}
		w := &bytes.Buffer{}
		if err := b.Write(w); (err != nil) != bb.wantErr {
			b.Errorf("%q. Bar.Write() error = %v, wantErr %v", tt.name, err, bb.wantErr)
			continue
		}
		if gotW := w.String(); gotW != bb.wantW {
			b.Errorf("%q. Bar.Write() = %v, want %v", tt.name, gotW, bb.wantW)
		}
	}
}

func TestWrite(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		w := &bytes.Buffer{}
		if err := Write(w, tt.args.data); (err != nil) != tt.wantErr {
			t.Errorf("%q. Write() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if gotW := w.String(); gotW != tt.wantW {
			t.Errorf("%q. Write() = %v, want %v", tt.name, gotW, tt.wantW)
		}
	}
}

func BenchmarkWrite(b *testing.B) {
	type args struct {
		data string
	}
	benchmarks := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		w := &bytes.Buffer{}
		if err := Write(w, tt.args.data); (err != nil) != bb.wantErr {
			b.Errorf("%q. Write() error = %v, wantErr %v", tt.name, err, bb.wantErr)
			continue
		}
		if gotW := w.String(); gotW != bb.wantW {
			b.Errorf("%q. Write() = %v, want %v", tt.name, gotW, bb.wantW)
		}
	}
}

func TestMultiWrite(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantW1  string
		wantW2  string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		w1 := &bytes.Buffer{}
		w2 := &bytes.Buffer{}
		got, got1, err := MultiWrite(w1, w2, tt.args.data)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. MultiWrite() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. MultiWrite() got = %v, want %v", tt.name, got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("%q. MultiWrite() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
		if gotW1 := w1.String(); gotW1 != tt.wantW1 {
			t.Errorf("%q. MultiWrite() gotW1 = %v, want %v", tt.name, gotW1, tt.wantW1)
		}
		if gotW2 := w2.String(); gotW2 != tt.wantW2 {
			t.Errorf("%q. MultiWrite() gotW2 = %v, want %v", tt.name, gotW2, tt.wantW2)
		}
	}
}

func BenchmarkMultiWrite(b *testing.B) {
	type args struct {
		data string
	}
	benchmarks := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantW1  string
		wantW2  string
		wantErr bool
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		w1 := &bytes.Buffer{}
		w2 := &bytes.Buffer{}
		got, got1, err := MultiWrite(w1, w2, tt.args.data)
		if (err != nil) != bb.wantErr {
			b.Errorf("%q. MultiWrite() error = %v, wantErr %v", tt.name, err, bb.wantErr)
			continue
		}
		if got != bb.want {
			b.Errorf("%q. MultiWrite() got = %v, want %v", tt.name, got, bb.want)
		}
		if got1 != bb.want1 {
			b.Errorf("%q. MultiWrite() got1 = %v, want %v", tt.name, got1, bb.want1)
		}
		if gotW1 := w1.String(); gotW1 != bb.wantW1 {
			b.Errorf("%q. MultiWrite() gotW1 = %v, want %v", tt.name, gotW1, bb.wantW1)
		}
		if gotW2 := w2.String(); gotW2 != bb.wantW2 {
			b.Errorf("%q. MultiWrite() gotW2 = %v, want %v", tt.name, gotW2, bb.wantW2)
		}
	}
}
