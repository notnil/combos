package combos_test

import (
	"testing"

	"github.com/notnil/combos"
)

func TestCombinations(t *testing.T) {
	results := combos.New(7, 5)
	if len(results) != 21 {
		t.Fatalf("expected %d results but got %d", 21, len(results))
	}
}
