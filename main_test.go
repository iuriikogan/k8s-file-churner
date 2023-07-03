package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreateFile(t *testing.T) {
	fileSizeBytes := 1024
	fileIndex := 1
	done := make(chan bool)

	go createFile(fileSizeBytes, fileIndex, done)

	select {
	case <-done:
		// File creation completed
		filePath := fmt.Sprintf("data/test_file%d.txt", fileIndex)
		_, err := os.Stat(filePath)
		if err != nil {
			t.Errorf("File '%s' does not exist: %s", filePath, err)
		}

		err = os.Remove(filePath)
		if err != nil {
			t.Errorf("Failed to delete file '%s': %s", filePath, err)
		}
	case <-time.After(5 * time.Second):
		t.Errorf("Timeout: File creation took too long")
	}
}

func TestChurnFiles(t *testing.T) {
	// Create test files
	err := os.WriteFile("data/test_file1.txt", []byte("Test file 1"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("data/test_file2.txt", []byte("Test file 2"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	err = os.WriteFile("data/not_test_file.txt", []byte("Not a test file"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	fileSizeBytes := 1024
	churnPercentage := 0.5
	done := make(chan bool)

	churnFiles(churnPercentage, fileSizeBytes, done)

	select {
	case <-done:
		// Verify that files starting with "test" are churned
		files, err := os.ReadDir("data")
		if err != nil {
			t.Fatalf("Failed to read directory: %s", err)
		}

		for _, file := range files {
			if file.Name() == "test_file1.txt" || file.Name() == "test_file2.txt" {
				t.Errorf("File '%s' should have been churned", file.Name())
			}
		}
	case <-time.After(5 * time.Second):
		t.Errorf("Timeout: Churn operation took too long")
	}

	// Clean up test files
	err = os.Remove(filepath.Join("data", "test_file1.txt"))
	if err != nil {
		t.Fatalf("Failed to delete file 'test_file1.txt': %s", err)
	}
	err = os.Remove(filepath.Join("data", "test_file2.txt"))
	if err != nil {
		t.Fatalf("Failed to delete file 'test_file2.txt': %s", err)
	}
	err = os.Remove(filepath.Join("data", "not_test_file.txt"))
	if err != nil {
		t.Fatalf("Failed to delete file 'not_test_file.txt': %s", err)
	}
}

func TestMain(m *testing.M) {
	// Set up test environment
	err := os.Mkdir("data", 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll("data")

	// Run tests
	code := m.Run()

	// Clean up test environment

	os.Exit(code)
}
