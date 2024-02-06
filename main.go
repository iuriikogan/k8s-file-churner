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
	// typecast ChurnIntervalMinutes(int64) to time.Duration to print it in minutes
	//
<<<<<<< HEAD
	var churnInterval = time.Duration(cfg.ChurnIntervalMinutes)
	log.Printf("Churn interval in minutes: %v\n", churnInterval)
	// log stuff
	log.Println("Starting K8s File Churner...\nAll testfiles will be written to tmp/testfiles directory")
	log.Printf("Size of each file in Mb: %v\n", cfg.SizeOfFileMB)
	log.Printf("Size of PVC in Gb: %v\n", cfg.SizeOfPVCGB)
	log.Printf("Churn percentage: %v\n", (cfg.ChurnPercentage))
=======
	churnInterval := time.Duration(cfg.ChurnIntervalMinutes) * time.Minute
	//
	log.Printf("Churn interval in minutes: %v\n", churnInterval)
	// log stuff
	log.Println("Starting K8s File Churner...\nAll testfiles will be written to app/testfiles directory")
	log.Printf("Size of each file in Mb: %d\n", cfg.SizeOfFileMB)
	log.Printf("Size of PVC in Gb: %d\n", cfg.SizeOfPVCGB)
	log.Printf("Churn percentage: %v\n", (cfg.ChurnPercentage * 100))
>>>>>>> parent of 2305d96 (something broken)
	log.Printf("Churn interval in minutes: %v\n", churnInterval)
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
	c := cron.New()       // create a new cron to log every 2 minute to ensure go routines are still running
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
	churnTicker := time.NewTicker(churnInterval) // create a ticker to churn files every churnInterval
	go func() {
<<<<<<< HEAD
		log.Printf("Churning %v percent of files every %v for %v", (cfg.ChurnPercentage * 100), churnInterval, cfg.ChurnDurationHours)
		startChurn := time.Now()
		done := make(chan bool)
=======
		log.Printf("Churning %v percent of files every %v", (cfg.ChurnPercentage * 100), churnInterval)

>>>>>>> parent of 2305d96 (something broken)
		for {
			select {
			case <-churnTicker.C: // churn every churnInterval
				utils.ChurnFiles(cfg.ChurnPercentage, fileSizeBytes, &wg)
			case <-time.After(120 * time.Second): // log every 120 seconds
				log.Println("Waiting to churn files")
<<<<<<< HEAD
			case <-time.After(cfg.ChurnDurationHours * time.Hour): // Stop churning after churnDurationHours
				defer churnTicker.Stop()
				done <- true
				log.Printf("Churned %v percent of files every %v for %v\nTook %s", (cfg.ChurnPercentage * 100), churnInterval, cfg.ChurnDurationHours, time.Since(startChurn))
=======
>>>>>>> parent of 2305d96 (something broken)
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
