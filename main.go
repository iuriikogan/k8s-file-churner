package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gogits/cron"
	"github.com/iuriikogan/k8s-file-churner/config"
	_ "go.uber.org/automaxprocs"
)

func main() {
	start := time.Now()
	// // log to custom file
	// logFile, err := os.OpenFile("var/log/k8sfilechurner.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.SetOutput(logFile)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Create the app/testfiles directory
	err = os.MkdirAll("app/testfiles", 0755)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("************************************\nK8s File Churner was made by Iurii Kogan - koganiurii@gmail.com \n************************************\n")
	log.Println("Starting K8s File Churner...\nAll testfiles will be written to app/testfiles directory")

	// fmt.Printf("Size of each file in Mb: %d\n", cfg.SizeOfFileMB)
	// fmt.Printf("Size of PVC in Gb: %d\n", cfg.SizeOfPVCGB)

	log.Printf("Size of each file in Mb: %d\n", cfg.SizeOfFileMB)
	log.Printf("Size of PVC in Gb: %d\n", cfg.SizeOfPVCGB)

	sizeOfPVCMB := int(cfg.SizeOfPVCGB * 999)             // convert size of PVC to MB
	numberOfFiles := ((sizeOfPVCMB) / (cfg.SizeOfFileMB)) // convert size of PVC to MB to calculate number of files to create
	log.Printf("Number of files to create: %d\n", numberOfFiles)
	fileSizeBytes := int(cfg.SizeOfFileMB * 1024 * 1024) // Convert file size from MB to bytes and convert to int
	var wg sync.WaitGroup
	wg.Add(numberOfFiles) // increment the wait group counter to numberoffiles to be created
	c := cron.New()       // create a new cron to log every 1 minute to ensure go routines are still running
	c.AddFunc("createFilesCron", "@every 1m", func() {
		log.Println("waiting for files to be created")
	})
	c.Start()
	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go createFile(fileSizeBytes, i, &wg)
	}
	// Wait for all the goroutines to finish
	wg.Wait()
	c.Stop() // stop the log cron

	live() // set the live probe

	log.Printf("Created %v files of size %vMb\nTook %s\n", numberOfFiles, cfg.SizeOfFileMB, time.Since(start))

	churnInterval := time.Duration(cfg.ChurnIntervalMinutes) // typecast ChurnIntervalMinutes to time.Duration

	churnTicker := time.NewTicker(churnInterval) // create a ticker to churn files every churnInterval
	go func() {
		log.Printf("Churning %v percent of files every %v", (cfg.ChurnPercentage * 100), churnInterval)

		for {
			select {
			case <-churnTicker.C: // churn every churnInterval
				churnFiles(cfg.ChurnPercentage, fileSizeBytes, &wg)
			case <-time.After(60 * time.Second): // log every 60 seconds
				log.Println("Waiting to churn files")
			}
		}
	}()

	// Keep the program running until interrupted
	<-make(chan struct{})
}

func createFile(fileSizeBytes int, fileIndex int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Generate a file name
	fileName := fmt.Sprintf("app/testfiles/%d.bin", fileIndex)
	file, err := os.Create(fileName) // Create the file
	if err != nil {
		unlive()
		log.Printf("Failed to create file '%s': %s\n", fileName, err)
		return
	}
	defer file.Close()

	writeRandomData(file, fileSizeBytes)
}

func writeRandomData(file *os.File, fileSizeBytes int) {
	chunkSize := 4096
	chunks := fileSizeBytes / chunkSize
	for i := 0; i < chunks; i++ {
		data := make([]byte, chunkSize) // create chunks of size 4096 bytes
		rand.Read(data)
		file.Write(data) // write the chunk to the file
	}
	remainingBytes := fileSizeBytes % chunkSize // calc the remaining bytes and keep looping through the remainder writing a chunk to the file each timen until remainingBytes !>0
	if remainingBytes > 0 {
		data := make([]byte, remainingBytes)
		rand.Read(data)
		file.Write(data)
	}
}

func churnFiles(churnPercentage float64, fileSizeBytes int, wg *sync.WaitGroup) {
	files, err := os.ReadDir("app/testfiles/") // read all
	if err != nil {
		unlive()
		log.Fatal(err)
		return
	}

	numberOfFiles := len(files)
	numberOfFilesToDelete := int(float64(numberOfFiles) * churnPercentage)
	if numberOfFilesToDelete == 0 {
		unlive()
		log.Println("No files to churn.")
		return
	}

	sort.Slice(files, func(i, j int) bool {
		fileNum1 := extractFileNumber(files[i].Name())
		fileNum2 := extractFileNumber(files[j].Name())
		return fileNum1 < fileNum2
	})

	// Delete the first numFilesToDelete files if they start with "test" and are not directories
	for i := 0; i < numberOfFilesToDelete; i++ {
		file := files[i]
		filePath := filepath.Join("app/testfiles", file.Name())
		err := os.Remove(filePath)
		if err != nil {
			log.Printf("Failed to delete file '%s': %s\n", filePath, err)
			continue
		}
		log.Printf("Deleted file '%s'\n", filePath)
	}

	wg.Add(numberOfFilesToDelete) // increment the wait group counter

	// Create the same number of files that were deleted in the sorted order
	for i := 0; i < numberOfFilesToDelete; i++ {
		log.Printf("Creating file app/testfiles/%d.bin\n", i)
		go createFile(fileSizeBytes, i, wg) //create files calls wg.done each
	}
}

// TODO helper functionsmove to utils package

func extractFileNumber(fileName string) int {
	// Extract the numeric part of the file name, assuming the format "app/testfiles/{number}.bin"
	numberStr := strings.TrimSuffix(strings.TrimPrefix(fileName, "app/testfiles/"), ".bin")
	fileNum, _ := strconv.Atoi(numberStr)
	return fileNum
}

func live() {
	os.WriteFile("tmp/healthy", []byte("ok"), 0664)
}

func unlive() {
	os.Remove("tmp/healthy")
}
