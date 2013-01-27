package ant

import (
	"bytes"
	"testing"
)

func TestHardResetToBytes(t *testing.T) {
	msg := HardReset()

	expected := []byte{
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		0x00, 0x00,
		0x00,
	}

	actual := msg.ToBytes()

	if len(expected) != len(actual) {
		t.Errorf("ToBytes result length incorrect. Expected %v, but actual was %v", len(expected), len(actual))
	}

	if !bytes.Equal(actual, expected) {
		t.Errorf("ToBytes result didn't equal expected 15 zero bytes")
	}
}
