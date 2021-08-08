package main

import (
	"fmt"

	"github.com/petershen0307/PEScanner/cmd"
	"github.com/petershen0307/PEScanner/models"
	singleworker "github.com/petershen0307/PEScanner/singleWorker"
)

func main() {
	config := cmd.Parse()
	fmt.Println(config)
	if config.Mode == models.Single {
		singleworker.Inventory(config.EntryFolder)
		singleworker.WriteReport(config.OutputDir)
	}
}
