package antport

import(
	"testing"
)

func TestConnect(t *testing.T) {
	ctx := NewContext()
	connected := Connect(ctx)

	if(!connected) {
		t.Error("Hardware is missing")
	}
}