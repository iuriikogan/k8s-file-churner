package test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCreateFile(t *testing.T) {
	// make sure number of files created is equal to number of files specified
	files, err := os.ReadDir("data/") // read the data dir and pull all files into files slice
	if len(files) != numberOfFiles {  // pull numberOfFiles from config
		t.Errorf("Number of files created does not match number of files specified. Expected %d, got %d", numberOfFiles, len(files))
	}
}

func TestWriteRandomData(t *testing.T) {
	fmt.Println("Test write random data") // TODO test writeRandomData
}

func TestChurnFiles(t *testing.T) {
	fmt.Println("Test delete files") // TODO test churnFiles
}
