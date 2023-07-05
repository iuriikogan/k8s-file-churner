package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/iuriikogan/k8s-file-churner/config"
)

func main() {
	// Create the data directory if it doesn't exist
	err := os.MkdirAll("testfiles", 0777)
	if err != nil {
		panic(err)
	} // panic if the directory cannot be created

	runtime.GOMAXPROCS(10) // set the number of threads to run
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	start := time.Now() // start the timer
	fmt.Printf("Size of each file in Mb: %d\n", cfg.SizeOfFileMB)
	fmt.Printf("Size of PVC in Gb: %d\n", cfg.SizeOfPVCGB)

	sizeOfPVCMB := cfg.SizeOfPVCGB * 1023
	numberOfFiles := (sizeOfPVCMB) / (cfg.SizeOfFileMB) // convert size of PVC to MB to calculate number of files to create
	fmt.Printf("Number of files to create: %d\n", numberOfFiles)

	fileSizeBytes := int(cfg.SizeOfFileMB * 1024 * 1024) // Convert file size from MB to bytes and convert to int
	fmt.Printf("Size of each file: %dMb\n", cfg.SizeOfFileMB)
	var wg sync.WaitGroup
	wg.Add(numberOfFiles) // increment the wait group counter

	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go createFile(fileSizeBytes, i, &wg)
	}

	// Wait for all the goroutines to finish
	wg.Wait()

	fmt.Printf("Created %v files of size %vMb\nTook %s\n", numberOfFiles, cfg.SizeOfFileMB, time.Since(start))

	churnInterval := time.Duration(cfg.ChurnIntervalMinutes * 60 * 1000 * 1000 * 1000)
	fmt.Printf("Churn interval: %v\n", churnInterval)

	churnTicker := time.NewTicker(churnInterval)
	go func() {
		log.Printf("Churning %v percent of files every %v", (cfg.ChurnPercentage * 100), churnInterval)

		for {
			select {
			case <-churnTicker.C:
				churnFiles(cfg.ChurnPercentage, fileSizeBytes, &wg)
			case <-time.After(10 * time.Second):
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
	fileName := fmt.Sprintf("testfiles/%d.txt", fileIndex)
	file, err := os.Create(fileName) // Create the file
	if err != nil {
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
		data := make([]byte, chunkSize)
		rand.Read(data)
		file.Write(data)
	}

	remainingBytes := fileSizeBytes % chunkSize
	if remainingBytes > 0 {
		data := make([]byte, remainingBytes)
		rand.Read(data)
		file.Write(data)
	}
}

func churnFiles(churnPercentage float64, fileSizeBytes int, wg *sync.WaitGroup) {
	files, err := os.ReadDir("testfiles/")
	if err != nil {
		log.Fatal(err)
		return
	}

	numberOfFiles := len(files)
	numberOfFilesToDelete := int(float64(numberOfFiles) * churnPercentage)
	if numberOfFilesToDelete == 0 {
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
		filePath := filepath.Join("testfiles", file.Name())
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
		log.Printf("Creating file testfiles/%d.txt\n", i)
		go createFile(fileSizeBytes, i, wg)
	}
}

func extractFileNumber(fileName string) int {
	// Extract the numeric part of the file name, assuming the format "testfiles/{number}.txt"
	numberStr := strings.TrimSuffix(strings.TrimPrefix(fileName, "testfiles/"), ".txt")
	fileNum, _ := strconv.Atoi(numberStr)
	return fileNum
}

// func live() {
// 	os.WriteFile("tmp/healthy", []byte("ok"), 0664) // liveness prob	if err != nil
// 	panic("unable to set liveness probe")
// }

// func unlive() {
// 	os.Remove("tmp/healthy") // liveness probe
// 	panic("unable to unset liveness probe")
// }
