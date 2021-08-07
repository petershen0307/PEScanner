package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/petershen0307/PEScanner/models"
)

func Parse() models.Config {
	config := models.Config{}
	currentDir, _ := os.Getwd()
	flag.IntVar((*int)(&(config.Mode)), "mode", (int)(models.Single), fmt.Sprintf("%v:single, %v:concurrent", models.Single, models.Concurrent))
	flag.IntVar((*int)(&(config.ConcurrentNumber)), "thread", models.MinConcurrentNumber, fmt.Sprintf("set concurrent number(%v-%v) in concurrent mode", models.MinConcurrentNumber, models.MaxConcurrentNumber))
	flag.StringVar(&config.EntryFolder, "entry", currentDir, "set scan start dir")
	flag.StringVar(&config.OutputDir, "output", currentDir, "set output dir")
	flag.Parse()

	checkAndCorrectConfigValue(&config)
	return config
}

func checkAndCorrectConfigValue(config *models.Config) {
	// check thread
	if config.ConcurrentNumber < models.MinConcurrentNumber || config.ConcurrentNumber > models.MaxConcurrentNumber {
		config.ConcurrentNumber = models.MinConcurrentNumber
		log.Println("invalid thread number set to default ", config.ConcurrentNumber)
	}
	// check mode
	if config.Mode != models.Single && config.Mode != models.Concurrent {
		config.Mode = models.Single
		log.Println("fix mode to Single ", config.Mode)
	}
	// check dir exist
	correctInputDir := func(dir *string) {
		currentDir, _ := os.Getwd()
		if fileInfo, err := os.Stat(*dir); !os.IsNotExist(err) {
			if !fileInfo.IsDir() {
				*dir = currentDir
				log.Println("input dir is not a dir. fix to current dir ", *dir)
			}
		} else {
			*dir = currentDir
			log.Println("fix input dir to current dir ", *dir)
		}
	}
	correctInputDir(&config.EntryFolder)
	correctInputDir(&config.OutputDir)
}
