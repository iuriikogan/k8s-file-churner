package main

import (
	"fmt"
	"github.com/iuriikogan/k8s-file-churner/utils"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(10) // set the number of threads to run
	// TODO liveness probe / readiness probe for the pod
	// TODO load config from ENV variables
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("failed to load the config", err)
	}
	return
	fmt.Printf("Size of each file in Mb: %d\n", config.sizeOfFileMB)
	fmt.Printf("Size of PVC in Gb: %d\n", config.sizeOfPVCGB)
	sizeOfPVCMB := config.sizeOfPVCGB * 1024
	numberOfFiles := (config.sizeOfPVCMB) / (config.sizeOfFileMB - 1) // convert size of PVC to MB to calculate number of files to create
	fmt.Printf("Number of files to create: %d\n", numberOfFiles)
	fileSizeBytes := int(sizeOfFileMB * 1024 * 1024) // Convert file size from MB to bytes and convert to int
	fmt.Printf("Size of each file: %dMb\n", sizeOfFileMB)
	done := make(chan bool) // which sets done when a createfile routine is created and closes the channel when done is true

	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go createFile(fileSizeBytes, i, done)
	}
	// Wait for all the goroutines to finish
	for i := 0; i < numberOfFiles; i++ {
		<-done // while done is true
	}
	fmt.Printf("created %v files of size %vMb\n Took %s\n", numberOfFiles, sizeOfFileMB, time.Since(start))
	// churnInterval := 30 * time.Second // int Churn interval seconds load from config
	// churnPercentage := 0.5            // float64 Churn percentage
	churnTicker := time.NewTicker(config.churnIntervalMinutes)
	go func() {
		log.Printf("Churning %v percent of files every %v", (config.churnPercentage * 100), churnInterval)
		for {
			select {
			case <-churnTicker.C:
				churnFiles(churnPercentage, fileSizeBytes, done)
			case <-time.After(10 * time.Second):
				log.Println("Waiting to churn files")
			}
		}
	}()

	// Keep the program running until interrupted
	<-make(chan struct{})
}

func createFile(fileSizeBytes int, fileIndex int, done chan<- bool) {
	// Generate a file name
	fileName := fmt.Sprintf("data/test_file%d.txt", fileIndex)
	// TODO check the directory exists and create it if it doesn't (currently done as part of the dockerfile)
	file, err := os.Create(fileName) // Create the file
	if err != nil {
		done <- false
		panic(err)
	}
	defer file.Close()

	// Write random data to the file and send a message to the channel when done
	writeRandomData(file, fileSizeBytes, err)
	if err != nil {
		done <- false
		panic(err)
	}
	log.Printf("Created file '%s' of size %vMb", file.Name(), int32(fileSizeBytes/1024/1024)) // TODO display filesize in GB
	done <- true
}

// write random data using the math/rand package since there is no need for crypotographically secure random data
func writeRandomData(file *os.File, fileSizeBytes int, err error) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	if err != nil {
		log.Printf("Failed to write data to file %s\n, Error: %s", file.Name(), err)
		panic(err)
	}

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

func churnFiles(churnPercentage float64, fileSizeBytes int, done chan<- bool) {
	files, err := os.ReadDir("data/")
	if err != nil {
		log.Fatal(err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	numberOfFiles := len(files)
	numberOfFilesToDelete := int(float64(numberOfFiles) * churnPercentage)
	if numberOfFilesToDelete == 0 {
		log.Println("No files to churn.")
		return
	}
	// Delete the first numFilesToDelete files if they start with "test" and are not directories
	for i := 0; i < numberOfFilesToDelete; i++ {
		file := files[i]
		done := make(chan bool)
		if strings.HasPrefix(file.Name(), "test") && !file.IsDir() { // check if file begins with test
			filePath := filepath.Join("data", file.Name())
			err := os.Remove(filePath)
			if err != nil {
				log.Printf("Failed to delete file '%s': %s\n", filePath, err)
				continue
			}
			log.Printf("Deleted file '%s'\n", filePath)
			go createFile(fileSizeBytes, i, done)
		}
		for i := 0; i < numberOfFiles; i++ {
		}
		<-done
	}
	done <- true
}
