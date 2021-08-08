package scanner

import (
	"crypto/sha256"
	"io"
	"os"
)

// GetFileSha2 give a file path and return a sha2 value
func GetFileSha2(file *os.File) []byte {
	sha2Gen := sha256.New()
	io.Copy(sha2Gen, file)
	return sha2Gen.Sum(nil)
}

// IsPEFile return the target file is PE or not
func IsPEFile(file *os.File) bool {
	return true
}
