package utils

import (
	"os"
	"sync"
	"testing"
)

func TestCreateFileWithRandomData(t *testing.T) {
	os.Mkdir("tmp/testfiles/", 0755)
	wg := sync.WaitGroup{}
	wg.Add(1)
	// Call the writeRandomData function
	CreateFileWithRandomData((1024 * 1024), 1, &wg)
	file, err := os.OpenFile("tmp/testfiles/1.bin", os.O_RDONLY, 0644)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	// Check if the file size matches the expected size
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}
	fileSize := fileInfo.Size()
	expectedSize := int64((1024 * 1024))
	if fileSize != expectedSize {
		t.Errorf("Expected file size %d, but got %d", expectedSize, fileSize)
	}
	cleanUpFiles()
}

func TestChurnFiles(t *testing.T) {
	os.Mkdir("tmp/testfiles/", 0755)
	// Create some test files in the directory
	for i := 0; i < 5; i++ {
		file, err := os.CreateTemp("tmp/testfiles/", "*.bin")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer file.Close()
	}
	// Get the initial number of files
	initialFiles, err := os.ReadDir("tmp/testfiles/")
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}
	initialFileCount := len(initialFiles)
	// Create a wait group
	var wg sync.WaitGroup
	// Call the churnFiles function
	ChurnFiles(0.2, (1024 * 1024), &wg)
	// Wait for the goroutines to finish
	wg.Wait()
	// Get the number of files after the churn operation
	updatedFiles, err := os.ReadDir("tmp/testfiles/")
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}
	updatedFileCount := len(updatedFiles)

	// Check the number of files in the directory
	expectedFiles := initialFileCount
	if updatedFileCount != expectedFiles {
		t.Errorf("Expected %d files, but got %d", expectedFiles, updatedFileCount)
	}
	cleanUpFiles()
}

func cleanUpFiles() {
	os.RemoveAll("tmp/testfiles/")
}
