package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)

}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("Didn't find seperator")
	}
	contentLengthBytes := header[len("Content-Length: "):] // "Content-Length: 15", this returns 15
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}
	_ = content
	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}
	return baseMessage.Method, content[:contentLength], nil
}

// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error) <- Scanner's SplitFunc Type, uses this internally
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}
	contentLengthBytes := header[len("Content-Length: "):]         // "Content-Length: 15", this returns 15
	contentLength, err := strconv.Atoi(string(contentLengthBytes)) // "Content-Length: 15", this returns 15
	if err != nil {
		return 0, nil, err
	}
	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength // +4 for the \r\n\r\n
	return totalLength, data[:totalLength], nil
}
