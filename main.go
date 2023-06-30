package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4) // set the number of threads to run
	// TODO load config from ENV variables
	// config, err := utils.LoadConfig("./")
	// if err != nil {
	// 	log.Fatal("failed to load the config", err)
	// }
	// return

	fileSizeGB := 1
	PVCSizeGB := 30
	numberOfFiles := (int(PVCSizeGB) / int(fileSizeGB))
	fileSizeBytes := int(fileSizeGB * 1024 * 1024 * 1023) // Convert file size from GB to bytes and convert to int
	done := make(chan bool)                               // which sets done when a createfile routine is created and closes the channel when done is true

	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go createFile(fileSizeBytes, i, done)
	}

	// Wait for all the goroutines to finish
	for i := 0; i < numberOfFiles; i++ {
		<-done // while done is true
	}

	churnInterval := 30 * time.Second // int Churn interval seconds load from config
	churnPercentage := 0.1            // float64 Churn percentage
	churnTicker := time.NewTicker(churnInterval)
	go func() {
		log.Printf("Churning %v percent of files every %v", (churnPercentage * 100), churnInterval)
		for {
			select {
			case <-churnTicker.C:
				churnFiles(churnPercentage, fileSizeBytes, done)
			case <-time.After(10 * time.Second):
				log.Println("waiting to churn files")
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
	file, err := os.Create(fileName) 	// Create the file
	if err != nil {
		log.Printf("Failed to create file '%s': %s\n", fileName, err)
		done <- false
		return
	}
	defer file.Close()

	// Write random data to the file and send a message to the channel when done
	writeRandomData(file, fileSizeBytes, err)
	if err != nil {
		log.Fatal(err)
		done <- false
		return
	}
	log.Printf("Created file '%s' of size %vGb", file.Name(), int(fileSizeBytes/1024/1024/1023)) // TODO display filesize in MB
	done <- true
}

func writeRandomData(file *os.File, fileSizeBytes int, err error) {
	if err != nil {
		log.Printf("Failed to write data to file %s\n, Error: %s", file.Name(), err)
	}
	data := make([]byte, int(fileSizeBytes)) // create of buffer of size fileSizeBytes
	file.Write(data)                         // write the buffer to the file
}

func churnFiles(churnPercentage float64, fileSizeBytes int, done chan<- bool) {
	files, err := os.ReadDir("data/") // read the data dir and pull all files into files slice
	if err != nil {
		log.Printf("Failed to read directory: %s\n", err)
		return
	}
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
		if strings.HasPrefix(file.Name(), "test") && !file.IsDir() { // TODO check why this is not working as expected
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
}
