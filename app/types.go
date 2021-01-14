package app

import (
	"crypto/sha256"
	"encoding/base64"
)

// File defines the structure of the file
type File struct {
	Name string
	Data []byte
}

// Hash calculates base64 encoded SHA256 digest of the Data field of the File
func (f File) Hash() string {
	h := sha256.New()
	h.Write(f.Data)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
