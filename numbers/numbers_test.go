package numbers

import (
	"testing"
)

func TestSign(t *testing.T) {
	result := Sign(2)
	if result != "+" {
		t.Fail()
	}
}
