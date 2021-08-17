package concurrentworker

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/petershen0307/PEScanner/models"
	"github.com/petershen0307/PEScanner/report"
	"github.com/petershen0307/PEScanner/scanner"
)

// filePathChan pass file path to scan worker
var filePathChan chan string

// reportChan pass models.Report to report worker
var reportChan chan models.Report

func init() {
	filePathChan = make(chan string, 100)
	reportChan = make(chan models.Report, 10)
}

func walkCallback(path string, d fs.DirEntry, err error) error {
	if !d.IsDir() {
		filePathChan <- path
		models.ProfilingMetric.ScanFiles++
	}
	return nil
}

func inventory(startDir string) {
	err := filepath.WalkDir(startDir, walkCallback)
	if err != nil {
		log.Println("filepath.WalkDir error:", err)
	}
}

func scanFileAndReport(filePath string) {
	file, openError := os.Open(filePath)
	if openError != nil {
		log.Printf("open file(%v) error(%v)\n", filePath, openError)
		return
	}
	defer file.Close()
	if scanner.IsPEFile(file) {
		oneReport := models.Report{FilePath: filePath, Sha2: scanner.GetFileSha2(file)}
		reportChan <- oneReport
	}
}

func scan(wg *sync.WaitGroup) {
	defer wg.Done()
	for filePath := range filePathChan {
		scanFileAndReport(filePath)
	}
}

func collectReport(outputDir string, wg *sync.WaitGroup) {
	defer wg.Done()
	reports := []models.Report{}
	for oneReport := range reportChan {
		reports = append(reports, oneReport)
	}
	report.Write(outputDir, models.Concurrent, reports)
	models.ProfilingMetric.PeFiles = len(reports)
}

func Run(config models.Config) {
	models.ProfilingMetric.Mode = models.Concurrent
	models.ProfilingMetric.StartTime = time.Now()
	scanWG := sync.WaitGroup{}
	for i := 0; i < config.ConcurrentNumber; i++ {
		go scan(&scanWG)
		scanWG.Add(1)
	}
	reportWG := sync.WaitGroup{}
	go collectReport(config.OutputDir, &reportWG)
	reportWG.Add(1)

	inventory(config.EntryFolder)
	// wait filePathChan is empty
	for len(filePathChan) != 0 {
		log.Println("Wait consume filePathChan")
		time.Sleep(1000 * time.Millisecond)
	}
	close(filePathChan)
	scanWG.Wait()
	// wait reportChan is empty
	for len(reportChan) != 0 {
		log.Println("Wait consume reportChan")
		time.Sleep(1000 * time.Millisecond)
	}
	close(reportChan)
	reportWG.Wait()
	models.ProfilingMetric.EndTime = time.Now()
	models.ProfilingMetric.Write(config.OutputDir)
}
