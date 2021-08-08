package scanner

import (
	"crypto/sha256"
	"encoding/binary"
	"io"
	"log"
	"os"
)

// GetFileSha2 give a file path and return a sha2 value
func GetFileSha2(file *os.File) []byte {
	sha2Gen := sha256.New()
	io.Copy(sha2Gen, file)
	return sha2Gen.Sum(nil)
}

const signatureOffset = 0x3C
const offsetSize = 4
const peSignatureValue = 0x4550

// IsPEFile return the target file is PE or not
func IsPEFile(file *os.File) bool {
	// check file size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Read file stat error", err)
		return false
	}
	// not PE file because file size not enough
	if fileInfo.Size() < signatureOffset {
		return false
	}
	signatureValueBytes := make([]byte, offsetSize)
	readN, err := file.ReadAt(signatureValueBytes, signatureOffset)
	if readN != offsetSize || (err != nil && err != io.EOF) {
		log.Printf("Read signature offset location(%v) not enough %v\n", readN, err)
		return false
	}
	sigLocationOffset := binary.LittleEndian.Uint32(signatureValueBytes)
	// not PE file because file size not enough
	if fileInfo.Size() < int64(sigLocationOffset) {
		return false
	}
	peSigBytes := make([]byte, offsetSize)
	readN, err = file.ReadAt(peSigBytes, int64(sigLocationOffset))
	if readN != offsetSize || err != nil {
		log.Printf("Read PE signature location(%v) not enough %v\n", readN, err)
		return false
	}
	return peSignatureValue == binary.LittleEndian.Uint32(peSigBytes)
}
