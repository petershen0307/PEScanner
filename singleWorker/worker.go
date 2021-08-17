package singleworker

import (
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
	models.ProfilingMetric.Mode = models.Single
	models.ProfilingMetric.StartTime = time.Now()
	inventory(config.EntryFolder)
	models.ProfilingMetric.EndTime = time.Now()
	models.ProfilingMetric.PeFiles = len(singleReports)
	report.Write(config.OutputDir, models.ProfilingMetric.Mode, singleReports)
	models.ProfilingMetric.Write(config.OutputDir)
}

func walkCallback(path string, d fs.DirEntry, err error) error {
	if !d.IsDir() {
		file, openError := os.Open(path)
		if openError != nil {
			log.Printf("open file(%v) error(%v)\n", path, openError)
			return nil
		}
		defer file.Close()
		models.ProfilingMetric.ScanFiles++
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
