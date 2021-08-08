package main

import (
	"fmt"

	"github.com/petershen0307/PEScanner/cmd"
	concurrentworker "github.com/petershen0307/PEScanner/concurrentWorker"
	"github.com/petershen0307/PEScanner/models"
	singleworker "github.com/petershen0307/PEScanner/singleWorker"
)

func main() {
	config := cmd.Parse()
	fmt.Println(config)
	switch config.Mode {
	case models.Single:
		singleworker.Run(config)
	case models.Concurrent:
		concurrentworker.Run(config)
	default:
	}
}
