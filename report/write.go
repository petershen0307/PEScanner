package report

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/petershen0307/PEScanner/models"
)

func Write(outputDir string, mode models.ScanMode, reports []models.Report) {
	filePath := filepath.Join(outputDir, fmt.Sprintf("%v-%v.json", mode.String(), time.Now().UTC().Unix()))
	log.Println("output file:", filePath)
	jsonBytes, err := json.MarshalIndent(reports, "", "    ")
	if err != nil {
		log.Println("json MarshalIndent failed:", err)
		return
	}
	os.WriteFile(filePath, jsonBytes, 0o755)
}

func WriteProfiling(outputDir string, mode models.ScanMode, start, end time.Time, scanFiles, peFiles int) {
	filePath := filepath.Join(outputDir, fmt.Sprintf("%v-%v.csv", mode.String(), time.Now().UTC().Unix()))
	header := "mode, start(MicroSec), end(MicroSec), execution(MicroSec), scanFiles, peFiles"
	values := fmt.Sprintf("%v,%v,%v,%v,%v,%v", mode.String(), start.UnixNano()/1000, end.UnixNano()/1000, end.Sub(start).Microseconds(), scanFiles, peFiles)
	outputStr := strings.Join([]string{header, values}, "\n")
	os.WriteFile(filePath, []byte(outputStr), 0o755)
}
