// Package shapes exercises tests generated in an external _test package.
package shapes_test

import (
	"bytes"
	"sync"
	"testing"

	shapes "github.com/cweill/gotests/testdata/packagetest/shapes"
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, tt.want, tt.c.F())
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
			assert.Equal(t, tt.want, r.Area())
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
			assert.Equal(t, tt.want, r.Label())
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
			assert.Equal(t, tt.want, b.Origin())
		})
	}
}

func TestNewRect(t *testing.T) {
	type args struct {
		a shapes.Point
		b shapes.Point
	}
	tests := []struct {
		name      string
		args      args
		want      *shapes.Rect
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shapes.NewRect(tt.args.a, tt.args.b)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWrite(t *testing.T) {
	type args struct {
		p shapes.Point
	}
	tests := []struct {
		name      string
		args      args
		wantW     string
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			tt.assertion(t, shapes.Write(w, tt.args.p))
			assert.Equal(t, tt.wantW, w.String())
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
			assert.Equal(t, tt.want, shapes.Sum(tt.args.ps...))
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
