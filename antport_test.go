package antport

import(
	"testing"
	"log"
)

func TestIsHardware(t *testing.T) {
	isPresent := IsHardwarePresent()

	log.Printf("Hardware is present: %s", isPresent)

	if(!isPresent) {
		t.Error("Hardware is missing")
	}
}
