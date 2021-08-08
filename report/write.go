package report

import (
	"encoding/json"
	"log"
	"os"

	"github.com/petershen0307/PEScanner/models"
)

func Write(filePath string, reports []models.Report) {
	jsonBytes, err := json.MarshalIndent(reports, "", "    ")
	if err != nil {
		log.Println("json MarshalIndent failed:", err)
		return
	}
	os.WriteFile(filePath, jsonBytes, 0o755)
}
