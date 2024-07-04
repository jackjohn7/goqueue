package lib

import (
	"bytes"
)

func Escape(raw []byte) []byte {
	return bytes.ReplaceAll(raw, []byte("\n"), []byte("\\n"))
}

// Remove escapings from string
func Unescape(raw []byte) []byte {
	return bytes.ReplaceAll(raw, []byte("\\n"), []byte("\n"))
}
