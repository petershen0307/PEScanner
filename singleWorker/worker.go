package singleworker

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/petershen0307/PEScanner/models"
	"github.com/petershen0307/PEScanner/report"
	"github.com/petershen0307/PEScanner/scanner"
)

// singleReports is for single worker report collection
var singleReports []models.Report

// initialize singleReports
func init() {
	singleReports = make([]models.Report, 0)
}

func Run(config models.Config) {
	inventory(config.EntryFolder)
	writeReport(config.OutputDir)
}

func walkCallback(path string, d fs.DirEntry, err error) error {
	if !d.IsDir() {
		file, openError := os.Open(path)
		if openError != nil {
			log.Printf("open file(%v) error(%v)\n", path, openError)
			return nil
		}
		defer file.Close()
		if scanner.IsPEFile(file) {
			singleReports = append(singleReports, models.Report{FilePath: path, Sha2: scanner.GetFileSha2(file)})
		}
	}
	return nil
}

func inventory(startDir string) {
	err := filepath.WalkDir(startDir, walkCallback)
	if err != nil {
		log.Println("filepath.WalkDir error:", err)
	}
}

func writeReport(reportDir string) {
	report.Write(filepath.Join(reportDir, fmt.Sprintf("single-%v.json", time.Now().UTC().Unix())), singleReports)
}
