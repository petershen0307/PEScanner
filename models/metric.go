package models

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Metric mature scan telemetry
type Metric struct {
	Mode      ScanMode
	ScanFiles int
	PeFiles   int
	StartTime time.Time
	EndTime   time.Time
}

var ProfilingMetric Metric

func (metric *Metric) Write(outputDir string) {
	filePath := filepath.Join(outputDir, fmt.Sprintf("%v-%v.csv", metric.Mode.String(), time.Now().UTC().Unix()))
	header := "mode, start(MicroSec), end(MicroSec), execution(MicroSec), scanFiles, peFiles"
	values := fmt.Sprintf("%v,%v,%v,%v,%v,%v", metric.Mode.String(), metric.StartTime.UnixNano()/1000, metric.EndTime.UnixNano()/1000, metric.EndTime.Sub(metric.StartTime).Microseconds(), metric.ScanFiles, metric.PeFiles)
	outputStr := strings.Join([]string{header, values}, "\n")
	os.WriteFile(filePath, []byte(outputStr), 0o755)
}
