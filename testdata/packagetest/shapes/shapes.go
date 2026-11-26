// Package shapes exercises tests generated in an external _test package.
package shapes

import (
	"io"
	"sync"
)

type Point struct{ X, Y int }

// Rect has an unexported field, which an external test cannot set.
type Rect struct {
	Min, Max Point
	label    string
}

// Box embeds a local and an imported type.
type Box struct {
	Point
	sync.Mutex
	Tags map[string][]Point
}

type Celsius float64

type config struct{ n int }

type server struct{}

func (c Celsius) F() float64            { return float64(c)*9/5 + 32 }
func (c *Celsius) Set(v float64)        { *c = Celsius(v) }
func (r Rect) Area() int                { return (r.Max.X - r.Min.X) * (r.Max.Y - r.Min.Y) }
func (r *Rect) Label() string           { return r.label }
func (b *Box) Origin() Point            { return b.Point }
func NewRect(a, b Point) (*Rect, error) { return &Rect{Min: a, Max: b}, nil }
func Write(w io.Writer, p Point) error  { _, err := io.WriteString(w, "p"); return err }
func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}
func Sum(ps ...Point) Point                     { return Point{} }
func Apply(f func(Point) Point, ch chan<- Rect) {}

// Not testable from an external package.
func Configure(c config) error { return nil }
func (s *server) Start() error { return nil }
func helper()                  {}
