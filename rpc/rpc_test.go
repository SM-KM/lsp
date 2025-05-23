package rpc

import (
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := EncodeMessage(EncodingExample{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := DecodeMessage([]byte(incomingMessage))
	contentLen := len(content)

	if err != nil {
		t.Fatal(err)
	}

	if contentLen != 15 {
		t.Fatalf("Expected: 15, Got: %d", contentLen)
	}

	if method != "hi" {
		t.Fatalf("Expected: hi, Got: %s", method)
	}
}
