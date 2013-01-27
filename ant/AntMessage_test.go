package ant

import (
	"bytes"
	"testing"
)

func TestAntMessageToBytes(t *testing.T) {

	sync := byte(0x01)
	id := byte(0x02)
	data := []byte{0x03, 0x04, 0x05}
	length := byte(len(data))

	expectedResultSize := 4 + len(data)

	msg := NewAntMessage(sync, id, data)
	result := msg.ToBytes()

	if result[0] != sync {
		t.Errorf("Sync byte not set correctly. Expected %v actual %v", sync, result[0])
	}

	if result[1] != length {
		t.Errorf("Length byte not set correctly. Expected %v actual %v", length, result[1])
	}

	if result[2] != id {
		t.Errorf("Id byte not set correctly. Expected %v actual %v", id, result[2])
	}

	if bytes.Equal(result[3:length], data) {
		t.Error("Data bytes not set correctly")
	}

	if len(result) != expectedResultSize {
		t.Error("Result array size is incorrect")
	}
}

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
	}
}
