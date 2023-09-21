package calc

import (
	"testing"
)

func TestMove(t *testing.T) {
	num := 10
	num = -num

	if num != -10 {
		t.Fatalf("nope")
	}
}
