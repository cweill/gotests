package goparser

import (
	"testing"

	"github.com/cweill/gotests/internal/models"
)

func TestQualifyTypes(t *testing.T) {
	local := map[string]bool{"Point": true, "Rect": true, "Set": true, "T": true, "config": true}
	expr := func(v string) *models.Expression { return &models.Expression{Value: v} }
	field := func(name, typ string) *models.Field { return &models.Field{Name: name, Type: expr(typ)} }

	tests := []struct {
		name       string
		fn         *models.Function
		wantReason string
		wantParams []string
		wantFields []string // Receiver field names that remain.
	}{
		{
			name: "composite types",
			fn: &models.Function{Name: "F", Parameters: []*models.Field{
				field("a", "[]Point"),
				field("b", "map[Point]*Rect"),
				field("c", "func(Point) (Rect, error)"),
				field("d", "chan<- Point"),
				field("e", "struct{ Point Point }"),
				field("f", "io.Reader"),
			}},
			wantParams: []string{
				"[]shapes.Point",
				"map[shapes.Point]*shapes.Rect",
				"func(shapes.Point) (shapes.Rect, error)",
				"chan<- shapes.Point",
				"struct{Point shapes.Point}",
				"io.Reader",
			},
		},
		{
			name:       "type parameters shadow local types",
			fn:         &models.Function{Name: "F", TypeParams: []*models.TypeParam{{Name: "T"}}, Parameters: []*models.Field{field("v", "Set[T]")}},
			wantParams: []string{"shapes.Set[T]"},
		},
		{
			name: "renamed receiver type parameters",
			fn: &models.Function{
				Name:       "Has",
				Receiver:   &models.Receiver{Field: field("s", "Set[T]")},
				Parameters: []*models.Field{field("v", "T")},
			},
			wantParams: []string{"T"},
		},
		{
			name: "unexported receiver fields are dropped",
			fn: &models.Function{
				Name: "Area",
				Receiver: &models.Receiver{
					Field:  field("r", "Rect"),
					Fields: []*models.Field{field("Min", "Point"), field("label", "string")},
				},
			},
			wantFields: []string{"Min"},
		},
		{
			name:       "unexported type in signature",
			fn:         &models.Function{Name: "Configure", Parameters: []*models.Field{field("c", "config")}},
			wantReason: "uses unexported type config",
		},
		{
			name:       "unexported receiver type",
			fn:         &models.Function{Name: "Start", Receiver: &models.Receiver{Field: field("s", "server")}},
			wantReason: "receiver type server is unexported",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QualifyTypes(tt.fn, "shapes", local); got != tt.wantReason {
				t.Fatalf("QualifyTypes() = %q, want %q", got, tt.wantReason)
			}
			for i, want := range tt.wantParams {
				if got := tt.fn.Parameters[i].Type.Value; got != want {
					t.Errorf("parameter %d = %q, want %q", i, got, want)
				}
			}
			if tt.fn.Receiver == nil || tt.wantFields == nil {
				return
			}
			if got := tt.fn.Receiver.Type.Value; got != "Rect" {
				t.Errorf("receiver type = %q, want it left unqualified", got)
			}
			var names []string
			for _, f := range tt.fn.Receiver.Fields {
				names = append(names, f.Name)
			}
			if len(names) != len(tt.wantFields) || names[0] != tt.wantFields[0] {
				t.Errorf("receiver fields = %v, want %v", names, tt.wantFields)
			}
		})
	}
}

// Fields declared together (a, b Point) share one Expression.
func TestQualifyTypesSharedExpression(t *testing.T) {
	shared := &models.Expression{Value: "Point"}
	fn := &models.Function{Name: "F", Parameters: []*models.Field{{Name: "a", Type: shared}, {Name: "b", Type: shared}}}
	QualifyTypes(fn, "shapes", map[string]bool{"Point": true})
	if shared.Value != "shapes.Point" {
		t.Errorf("shared type = %q, want %q", shared.Value, "shapes.Point")
	}
}
