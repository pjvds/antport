package antport

import (
	"testing"
)

func TestGenerateChecksum(t *testing.T) {
	data := []byte{
		178,
		164,
		59,
		176,
		204,
		75,
		97,
	}

	expected := byte(0)
	for i := 0; i < len(data); i++ {
		expected = expected ^ data[i]
	}

	actual := GenerateChecksum(data)

	if expected != actual {
		t.Errorf("checksum result (%v) was not as expected (%v)", actual, expected)
	} else {
		t.Log("checksum result correct")
	}
}
