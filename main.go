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
	runtime.GOMAXPROCS(10) // set the number of threads to run
	start := time.Now()    // Start the timer
	// config, err := utils.LoadConfig("./")
	// if err != nil {
	// 	log.Fatal("failed to load the config", err)
	// }
	// return
	sizeOfFileGB := 1
	sizeOfPVCGB := 30
	numberOfFiles := (int(sizeOfPVCGB) / int(sizeOfFileGB)) // TODO  Calculate the number of files to create based on PVC Size and File Size ENV Variables

	fileSizeBytes := sizeOfFileGB * 1024 * 1024 * 1023 // Convert file size from GB to bytes and convert to int

	done := make(chan bool) // which sets done when a createfile routine is created and closes the channel when done is true

	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go createFile(fileSizeBytes, i, done)
	}

	// Wait for all the goroutines to finish
	for i := 0; i < numberOfFiles; i++ {
		<-done // while done is true
	}

	churnInterval := 30 * time.Second // Churn interval
	churnPercentage := 0.2            // Churn percentage
	churnTicker := time.NewTicker(churnInterval)
	go func() {
		for {
			select {
			case <-churnTicker.C:
				churnFiles(churnPercentage, fileSizeBytes)
			case <-time.After(5 * time.Second):
				log.Printf("Time to next churn: %s", churnInterval-(time.Since(start)%churnInterval))
			}
		}
	}()

	// Keep the program running until interrupted
	<-make(chan struct{})
}

func createFile(fileSizeBytes int, fileIndex int, done chan<- bool) {
	// Generate a file name
	fileName := fmt.Sprintf("data/test_file%d.txt", fileIndex)
	// TODO check the directory exists and create it if it doesn't
	// Create the file
	file, err := os.Create(fileName)
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
	log.Printf("Created file '%s' of size %v", file.Name(), int(fileSizeBytes)) // TODO display filesize in MB
	done <- true
}

func writeRandomData(file *os.File, fileSizeBytes int, err error) {
	if err != nil {
		log.Printf("Failed to write data to file %s\n, Error: %s", file.Name(), err)
	}
	data := make([]byte, int(fileSizeBytes)) // create of buffer of size fileSizeBytes
	file.Write(data)                         // write the buffer to the file
}

func churnFiles(churnPercentage float64, fileSizeBytes int) {
	files, err := os.ReadDir("data/") // read the data dir and pull all files into files slice
	if err != nil {
		log.Printf("Failed to read directory: %s\n", err)
		return
	}
	numFiles := len(files)
	numFilesToDelete := int(float64(numFiles) * churnPercentage)
	if numFilesToDelete == 0 {
		log.Println("No files to churn.")
		return
	}
	// Delete the first numFilesToDelete files if they start with "test" and are not directories
	for i := 0; i < numFilesToDelete; i++ {
		file := files[i]
		if strings.HasPrefix(file.Name(), "test") && !file.IsDir() {
			filePath := filepath.Join("data", file.Name())
			err := os.Remove(filePath)
			if err != nil {
				log.Printf("Failed to delete file '%s': %s\n", filePath, err)
				continue
			}
			// launch a new createFile routine to replace the deleted file
			go createFile(fileSizeBytes, i, nil)
			log.Printf("File created '%s'", filePath)
		}
	}
}
