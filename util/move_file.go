package util

import (
	"fmt"
	"io"
	"os"
)

func MoveFile(sourcePath, destPath string) error {

	var err error

	inputFile, err := os.Open(sourcePath)

	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}

	defer inputFile.Close()

	outputFile, err := os.Create(destPath)

	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}

	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)

	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}

	err = os.Remove(sourcePath)

	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}

	return nil

}
