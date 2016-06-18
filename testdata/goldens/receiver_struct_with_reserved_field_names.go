package testdata

import "testing"

func TestReserved_DontFail(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Receiver fields.
		rName        string
		rBreak       string
		rDefault     string
		rFunc        string
		rInterface   string
		rSelect      string
		rCase        string
		rDefer       string
		rGo          string
		rMap         string
		rStruct      string
		rChan        string
		rElse        string
		rGoto        string
		rPackage     string
		rSwitch      string
		rConst       string
		rFallthrough string
		rIf          string
		rRange       string
		rType        string
		rContinue    string
		rFor         string
		rImport      string
		rReturn      string
		rVar         string
		// Expected results.
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		r := &Reserved{
			Name:        tt.rName,
			Break:       tt.rBreak,
			Default:     tt.rDefault,
			Func:        tt.rFunc,
			Interface:   tt.rInterface,
			Select:      tt.rSelect,
			Case:        tt.rCase,
			Defer:       tt.rDefer,
			Go:          tt.rGo,
			Map:         tt.rMap,
			Struct:      tt.rStruct,
			Chan:        tt.rChan,
			Else:        tt.rElse,
			Goto:        tt.rGoto,
			Package:     tt.rPackage,
			Switch:      tt.rSwitch,
			Const:       tt.rConst,
			Fallthrough: tt.rFallthrough,
			If:          tt.rIf,
			Range:       tt.rRange,
			Type:        tt.rType,
			Continue:    tt.rContinue,
			For:         tt.rFor,
			Import:      tt.rImport,
			Return:      tt.rReturn,
			Var:         tt.rVar,
		}
		if got := r.DontFail(); got != tt.want {
			t.Errorf("%q. Reserved.DontFail() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
