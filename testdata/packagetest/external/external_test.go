package external_test

import (
	"testing"

	"github.com/cweill/gotests/testdata/packagetest/external"
)

func TestExisting(t *testing.T) {
	if got := external.Existing(); got != 0 {
		t.Errorf("Existing() = %v, want 0", got)
	}
}
