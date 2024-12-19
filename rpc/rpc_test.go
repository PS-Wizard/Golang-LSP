package rpc_test

import (
	"educationallsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n {\"testing\":true}"
	got := rpc.EncodeMessage(EncodingExample{Testing: true})
	if expected == got {
		t.Fatalf("got %s, want %s", got, expected)
	}
}

func TestDecode(t *testing.T) {

	input := "Content-Length: 15\r\n\r\n{\"method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(input))
	got := len(content)
	want := 15
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("Got %d, Want %d", got, want)
	}
	if method != "hi" {
		t.Fatalf("Expected 'hi' got %s", method)
	}
}
