package testdata

import "testing"

func TestReserved_DontFail(t *testing.T) {
	type fields struct {
		Name        string
		Break       string
		Default     string
		Func        string
		Interface   string
		Select      string
		Case        string
		Defer       string
		Go          string
		Map         string
		Struct      string
		Chan        string
		Else        string
		Goto        string
		Package     string
		Switch      string
		Const       string
		Fallthrough string
		If          string
		Range       string
		Type        string
		Continue    string
		For         string
		Import      string
		Return      string
		Var         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := &Reserved{
			Name:        tt.fields.Name,
			Break:       tt.fields.Break,
			Default:     tt.fields.Default,
			Func:        tt.fields.Func,
			Interface:   tt.fields.Interface,
			Select:      tt.fields.Select,
			Case:        tt.fields.Case,
			Defer:       tt.fields.Defer,
			Go:          tt.fields.Go,
			Map:         tt.fields.Map,
			Struct:      tt.fields.Struct,
			Chan:        tt.fields.Chan,
			Else:        tt.fields.Else,
			Goto:        tt.fields.Goto,
			Package:     tt.fields.Package,
			Switch:      tt.fields.Switch,
			Const:       tt.fields.Const,
			Fallthrough: tt.fields.Fallthrough,
			If:          tt.fields.If,
			Range:       tt.fields.Range,
			Type:        tt.fields.Type,
			Continue:    tt.fields.Continue,
			For:         tt.fields.For,
			Import:      tt.fields.Import,
			Return:      tt.fields.Return,
			Var:         tt.fields.Var,
		}
		if got := r.DontFail(); got != tt.want {
			t.Errorf("%q. Reserved.DontFail() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func BenchmarkReserved_DontFail(b *testing.B) {
	type fields struct {
		Name        string
		Break       string
		Default     string
		Func        string
		Interface   string
		Select      string
		Case        string
		Defer       string
		Go          string
		Map         string
		Struct      string
		Chan        string
		Else        string
		Goto        string
		Package     string
		Switch      string
		Const       string
		Fallthrough string
		If          string
		Range       string
		Type        string
		Continue    string
		For         string
		Import      string
		Return      string
		Var         string
	}
	benchmarks := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add benchmark cases.
	}
	for _, bb := range benchmarks {
		r := &Reserved{
			Name:        bb.fields.Name,
			Break:       bb.fields.Break,
			Default:     bb.fields.Default,
			Func:        bb.fields.Func,
			Interface:   bb.fields.Interface,
			Select:      bb.fields.Select,
			Case:        bb.fields.Case,
			Defer:       bb.fields.Defer,
			Go:          bb.fields.Go,
			Map:         bb.fields.Map,
			Struct:      bb.fields.Struct,
			Chan:        bb.fields.Chan,
			Else:        bb.fields.Else,
			Goto:        bb.fields.Goto,
			Package:     bb.fields.Package,
			Switch:      bb.fields.Switch,
			Const:       bb.fields.Const,
			Fallthrough: bb.fields.Fallthrough,
			If:          bb.fields.If,
			Range:       bb.fields.Range,
			Type:        bb.fields.Type,
			Continue:    bb.fields.Continue,
			For:         bb.fields.For,
			Import:      bb.fields.Import,
			Return:      bb.fields.Return,
			Var:         bb.fields.Var,
		}
		if got := r.DontFail(); got != bb.want {
			b.Errorf("%q. Reserved.DontFail() = %v, want %v", tt.name, got, bb.want)
		}
	}
}
