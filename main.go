package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/gogits/cron"
	"github.com/iuriikogan/k8s-file-churner/config"
	"github.com/iuriikogan/k8s-file-churner/utils"
	_ "go.uber.org/automaxprocs"
)

func main() {
	// start the timer
	start := time.Now()
	// Load the config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Create the app/testfiles directory
	err = os.MkdirAll("app/testfiles", 0755)
	if err != nil {
		log.Fatal(err.Error())
	}
	//
	//
	// typecast ChurnIntervalMinutes to time.Duration
	//
	printChurnInterval := time.Duration.Minutes(cfg.ChurnIntervalMinutes)
	//
	log.Printf("Churn interval in minutes: %d\n", printChurnInterval)
	// log stuff
	log.Println("Starting K8s File Churner...\nAll testfiles will be written to app/testfiles directory")
	log.Printf("Size of each file in Mb: %d\n", cfg.SizeOfFileMB)
	log.Printf("Size of PVC in Gb: %d\n", cfg.SizeOfPVCGB)
	log.Printf("Churn percentage: %v\n", (cfg.ChurnPercentage * 100))
	log.Printf("Churn interval in minutes: %v\n", printChurnInterval)
	//
	//
	// calculate number of files to create
	//
	//
	sizeOfPVCMB := int(cfg.SizeOfPVCGB * 999)             // convert size of PVC to MB
	numberOfFiles := ((sizeOfPVCMB) / (cfg.SizeOfFileMB)) // convert size of PVC to MB to calculate number of files to create
	log.Printf("Number of files to create: %d\n", numberOfFiles)
	fileSizeBytes := int(cfg.SizeOfFileMB * 1024 * 1024) // Convert file size from MB to bytes and convert to int
	//
	//
	// start creating the files
	//
	//
	var wg sync.WaitGroup
	wg.Add(numberOfFiles) // increment the wait group counter to numberoffiles to be created
	c := cron.New()       // create a new cron to log every 1 minute to ensure go routines are still running
	c.AddFunc("createFilesCron", "@every 2m", func() {
		log.Println("waiting for files to be created")
	})
	c.Start()
	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go utils.CreateFileWithRandomData(fileSizeBytes, i, &wg)
	}
	// Wait for all the goroutines to finish
	wg.Wait()
	c.Stop() // stop the log cron
	// once all files are created, set the live probe
	utils.Live()
	// log the number of files created, their size and the time it took
	log.Printf("Created %v files of size %vMb\nTook %s\n", numberOfFiles, cfg.SizeOfFileMB, time.Since(start))
	//
	//
	// start churning the files
	//
	//
	churnTicker := time.NewTicker(cfg.ChurnIntervalMinutes) // create a ticker to churn files every churnInterval
	go func() {
		log.Printf("Churning %v percent of files every %v", (cfg.ChurnPercentage * 100), printChurnInterval)

		for {
			select {
			case <-churnTicker.C: // churn every churnInterval
				utils.ChurnFiles(cfg.ChurnPercentage, fileSizeBytes, &wg)
			case <-time.After(120 * time.Second): // log every 120 seconds
				log.Println("Waiting to churn files")
			}
		}
	}()
	//
	//
	// (this is a hack to keep the program running until interrupted)
	//
	//
	<-make(chan struct{})
}
