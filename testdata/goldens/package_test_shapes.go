// Package shapes exercises tests generated in an external _test package.
package shapes_test

import (
	"bytes"
	"reflect"
	"sync"
	"testing"

	shapes "github.com/cweill/gotests/testdata/packagetest/shapes"
)

func TestCelsius_F(t *testing.T) {
	tests := []struct {
		name string
		c    shapes.Celsius
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.F(); got != tt.want {
				t.Errorf("Celsius.F() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCelsius_Set(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		c    *shapes.Celsius
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Set(tt.args.v)
		})
	}
}

func TestRect_Area(t *testing.T) {
	type fields struct {
		Min shapes.Point
		Max shapes.Point
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := shapes.Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Rect.Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Label(t *testing.T) {
	type fields struct {
		Min shapes.Point
		Max shapes.Point
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &shapes.Rect{
				Min: tt.fields.Min,
				Max: tt.fields.Max,
			}
			if got := r.Label(); got != tt.want {
				t.Errorf("Rect.Label() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_Origin(t *testing.T) {
	type fields struct {
		Point shapes.Point
		Mutex sync.Mutex
		Tags  map[string][]shapes.Point
	}
	tests := []struct {
		name   string
		fields fields
		want   shapes.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &shapes.Box{
				Point: tt.fields.Point,
				Mutex: tt.fields.Mutex,
				Tags:  tt.fields.Tags,
			}
			if got := b.Origin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Box.Origin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRect(t *testing.T) {
	type args struct {
		a shapes.Point
		b shapes.Point
	}
	tests := []struct {
		name    string
		args    args
		want    *shapes.Rect
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shapes.NewRect(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewRect() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	type args struct {
		p shapes.Point
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
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := shapes.Write(w, tt.args.p); (err != nil) != tt.wantErr {
				t.Fatalf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("Write() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shapes.Max[int](tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		ps []shapes.Point
	}
	tests := []struct {
		name string
		args args
		want shapes.Point
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shapes.Sum(tt.args.ps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApply(t *testing.T) {
	type args struct {
		f  func(shapes.Point) shapes.Point
		ch chan<- shapes.Rect
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shapes.Apply(tt.args.f, tt.args.ch)
		})
	}
}
