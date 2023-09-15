package db

import (
	"os"
	"os/exec"
	"testing"
)

const (
	TEST_FILE_NAME = "test.g"
)

func TestWriteGFile(t *testing.T) {
	file, err := os.Create(TEST_FILE_NAME)
	//defer os.Remove(TEST_FILE_NAME)

	if err != nil {
		t.Fatalf("Could not create file: %s, details: %v", TEST_FILE_NAME, err)
	}

	if _, err := writeDb(file, "test", 1.0); err != nil {
		t.Fatalf("Could not write data to database file. details: %v", err)
	}

	mged := exec.Command("mged", "-c", TEST_FILE_NAME)
	mged.Stdout = os.Stdout

	if err = mged.Run(); err != nil {
		t.Fatalf("Execution of mged failed! details: %v", err)
	}
}
