package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/iuriikogan/k8s-file-churner/utils"
)

func main() {
	// Create the data directory if it doesn't exist
	err := os.MkdirAll("data", 0755)
	if err != nil {
		panic(err)
	} // panic if the directory cannot be created
	// TODO fix live() unlive() functions // k8s liveness probe
	runtime.GOMAXPROCS(4)             // set the number of threads to run
	config, err := utils.LoadConfig() // load the config from the current directory
	if err != nil {
		// unlive()
		log.Printf("Failed to load config: %v", err)
	}
	start := time.Now() // start the timer
	fmt.Printf("Size of each file in Mb: %d\n", config.SizeOfFileMB)

	fmt.Printf("Size of PVC in Gb: %d\n", config.SizeOfPVCGB)

	sizeOfPVCMB := config.SizeOfPVCGB * 1024

	numberOfFiles := (sizeOfPVCMB) / (config.SizeOfFileMB) // convert size of PVC to MB to calculate number of files to create

	fmt.Printf("Number of files to create: %d\n", numberOfFiles)

	fileSizeBytes := int(config.SizeOfFileMB * 1024 * 1024) // Convert file size from MB to bytes and convert to int
	fmt.Printf("Size of each file: %dMb\n", config.SizeOfFileMB)
	done := make(chan bool) // which sets done when a createfile routine is created and closes the channel when done is true

	// Launch a goroutine for each file creation
	for i := 0; i < numberOfFiles; i++ {
		go createFile(fileSizeBytes, i, done)
	}
	// Wait for all the goroutines to finish
	for i := 0; i < numberOfFiles; i++ {
		<-done // while done is true
	}
	// unlive() // set the liveness probe to false until the churn starts
	fmt.Printf("created %v files of size %vMb\n Took %s\n", numberOfFiles, config.SizeOfFileMB, time.Since(start))
	churnInterval := time.Duration(config.ChurnIntervalMinutes * 60 * 1000 * 1000 * 1000)
	fmt.Printf("Churn interval: %v\n", churnInterval)
	// time.Duration hurn Interval in Minutes due to type mismatch has to be resolved here instead of in the config.go file
	churnTicker := time.NewTicker(churnInterval)
	go func() {
		log.Printf("Churning %v percent of files every %v", (config.ChurnPercentage * 100), churnInterval)
		// live() // set the liveness probe to true
		for {
			select {
			case <-churnTicker.C:
				// live()
				churnFiles(config.ChurnPercentage, fileSizeBytes, done)
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
		// unlive()
		panic(err)
	}
	defer file.Close()

	// Write random data to the file and send a message to the channel when done
	writeRandomData(file, fileSizeBytes, err)
	if err != nil {
		done <- false
		// unlive()
		panic(err)
	}
	done <- true
}

// write random data using the math/rand package since there is no need for crypotographically secure random data
func writeRandomData(file *os.File, fileSizeBytes int, err error) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	if err != nil {
		log.Printf("Failed to write data to file %s\n, Error: %s", file.Name(), err)
		// unlive()
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
		// unlive()
		log.Fatal(err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	numberOfFiles := len(files)
	numberOfFilesToDelete := int(float64(numberOfFiles) * churnPercentage)
	if numberOfFilesToDelete == 0 {
		// unlive()
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

func live() {
	os.WriteFile("tmp/healthy", []byte("ok"), 0664) // liveness prob	if err != nil
	panic("unable to set liveness probe")
}

func unlive() {
	os.Remove("tmp/healthy") // liveness probe
	panic("unable to unset liveness probe")
}
