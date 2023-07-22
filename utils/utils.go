package utils

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
)

func CreateFileWithRandomData(fileSizeBytes int, fileIndex int, wg *sync.WaitGroup) (file *os.File) {
	defer wg.Done()
	// Generate a file name
	fileName := fmt.Sprintf("app/testfiles/%d.bin", fileIndex)
	file, err := os.Create(fileName) // Create the file
	if err != nil {
		Unlive()
		log.Printf("Failed to create file '%s': %s\n", fileName, err)
		return
	}
	defer file.Close()
	// Write random data to the file
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
	return file
}

func ChurnFiles(churnPercentage float64, fileSizeBytes int, wg *sync.WaitGroup) {
	files, err := os.ReadDir("app/testfiles/") // read all
	if err != nil {
		Unlive()
		log.Fatal(err)
		return
	}

	numberOfFiles := len(files)
	numberOfFilesToDelete := int(float64(numberOfFiles) * churnPercentage)
	if numberOfFilesToDelete == 0 {
		Unlive()
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

	wg.Add(numberOfFilesToDelete) // increment the wait group counter to the num of files to be deleted

	// Create the same number of files that were deleted in the sorted order
	for i := 0; i < numberOfFilesToDelete; i++ {
		log.Printf("Creating file app/testfiles/%d.bin\n", i)
		go CreateFileWithRandomData(fileSizeBytes, i, wg) //create files calls wg.done each iteration until = num of files to be deleted
	}
}

func extractFileNumber(fileName string) int {
	// Extract the numeric part of the file name, assuming the format "app/testfiles/{number}.bin"
	numberStr := strings.TrimSuffix(strings.TrimPrefix(fileName, "app/testfiles/"), ".bin")
	fileNum, _ := strconv.Atoi(numberStr)
	return fileNum
}

func Live() {
	os.WriteFile("tmp/healthy", []byte("ok"), 0664)
}

func Unlive() {
	os.Remove("tmp/healthy")
}
