package models

import (
	"encoding/hex"
	"encoding/json"
)

// Report structure collect file path and sha2 value
type Report struct {
	FilePath string
	Sha2     []byte
}

func (report Report) MarshalJSON() ([]byte, error) {
	alterReport := struct {
		FilePath string
		Sha2     string
	}{
		FilePath: report.FilePath,
		Sha2:     hex.EncodeToString(report.Sha2),
	}
	return json.Marshal(alterReport)
}
