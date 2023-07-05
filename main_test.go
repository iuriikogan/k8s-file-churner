package main

import (
	"os"
	"strconv"
	"sync"
	"testing"
)

func TestCreateFile(t *testing.T) {
	os.Mkdir("tmp", 0755)
	// Create a temporary file for testing
	file, err := os.CreateTemp("tmp", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Create a wait group
	var wg sync.WaitGroup

	// Add 1 to the wait group counter
	wg.Add(1)

	// Call the createFile function
	createFile(1024, 0, &wg)

	// Wait for the goroutine to finish
	wg.Wait()

	// Check if the file exists
	if _, err := os.Stat(file.Name()); os.IsNotExist(err) {
		t.Errorf("Expected file to be created, but it doesn't exist")
	}
}
func TestWriteRandomData(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("testfiles", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(file.Name())

	// Call the writeRandomData function
	writeRandomData(file, 1024)

	// Check if the file size matches the expected size
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()
	expectedSize := int64(1024)
	if fileSize != expectedSize {
		t.Errorf("Expected file size %d, but got %d", expectedSize, fileSize)
	}
}

func TestChurnFiles(t *testing.T) {
	// Create some test files in the directory
	for i := 0; i < 5; i++ {
		file, err := os.CreateTemp("testfiles", strconv.Itoa(i))
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer file.Close()
	}
	// Get the initial number of files
	initialFiles, err := os.ReadDir("testfiles")
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}
	initialFileCount := len(initialFiles)
	// Create a wait group
	var wg sync.WaitGroup
	// Call the churnFiles function
	churnFiles(0.2, 1024, &wg)
	// Wait for the goroutines to finish
	wg.Wait()
	// Get the number of files after the churn operation
	updatedFiles, err := os.ReadDir("testfiles")
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}
	updatedFileCount := len(updatedFiles)

	// Check the number of files in the directory
	expectedFiles := initialFileCount
	if updatedFileCount != expectedFiles {
		t.Errorf("Expected %d files, but got %d", expectedFiles, updatedFileCount)
	}
}
