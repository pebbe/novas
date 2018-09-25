package novas_test

import (
	"github.com/pebbe/novas"

	"fmt"
	"testing"
)

func TestJulian(t *testing.T) {
	time := novas.Julian(2458387.000000001)
	got := fmt.Sprintf("%.9f", time.ToJulian())
	exp := "2458387.000000001"
	if exp != got {
		t.Errorf("Expected %s, got %s", exp, got)
	}
}
