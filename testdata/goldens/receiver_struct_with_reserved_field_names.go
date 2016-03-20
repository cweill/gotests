package testdata

import "testing"

func TestReservedDontFail(t *testing.T) {
	tests := []struct {
		name         string
		fname        string
		fbreak       string
		fdefault     string
		ffunc        string
		finterface   string
		fselect      string
		fcase        string
		fdefer       string
		fgo          string
		fmap         string
		fstruct      string
		fchan        string
		felse        string
		fgoto        string
		fpackage     string
		fswitch      string
		fconst       string
		ffallthrough string
		fif          string
		frange       string
		ftype        string
		fcontinue    string
		ffor         string
		fimport      string
		freturn      string
		fvar         string
		want         string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := &Reserved{
			Name:        tt.fname,
			Break:       tt.fbreak,
			Default:     tt.fdefault,
			Func:        tt.ffunc,
			Interface:   tt.finterface,
			Select:      tt.fselect,
			Case:        tt.fcase,
			Defer:       tt.fdefer,
			Go:          tt.fgo,
			Map:         tt.fmap,
			Struct:      tt.fstruct,
			Chan:        tt.fchan,
			Else:        tt.felse,
			Goto:        tt.fgoto,
			Package:     tt.fpackage,
			Switch:      tt.fswitch,
			Const:       tt.fconst,
			Fallthrough: tt.ffallthrough,
			If:          tt.fif,
			Range:       tt.frange,
			Type:        tt.ftype,
			Continue:    tt.fcontinue,
			For:         tt.ffor,
			Import:      tt.fimport,
			Return:      tt.freturn,
			Var:         tt.fvar,
		}
		if got := r.DontFail(); got != tt.want {
			t.Errorf("%q. Reserved.DontFail() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
