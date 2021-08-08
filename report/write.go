package report

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/petershen0307/PEScanner/models"
)

func Write(filePath string, reports []models.Report, start, end time.Time) {
	log.Println("output file:", filePath)
	jsonBytes, err := json.MarshalIndent(struct {
		Reports             []models.Report
		StartTime           int64
		EndTime             int64
		DiffTimeMicroSecond int64
	}{
		Reports:             reports,
		StartTime:           start.UnixNano() / 1000,
		EndTime:             start.UnixNano() / 1000,
		DiffTimeMicroSecond: end.Sub(start).Microseconds(),
	}, "", "    ")
	if err != nil {
		log.Println("json MarshalIndent failed:", err)
		return
	}
	os.WriteFile(filePath, jsonBytes, 0o755)
}
