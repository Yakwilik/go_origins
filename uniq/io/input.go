package io

import (
	"bufio"
	"log"
	"os"
)

func readLines(file *os.File) (lines []string, err error) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func GetLines(inputFilename string) (lines []string, err error) {
	input := os.Stdin

	if inputFilename != "" {
		input, err = os.Open(inputFilename)
		if err != nil {
			return lines, err
		}
		defer func() {
			err := input.Close()
			if err != nil {
				log.Printf("error occured while closing file: %s", err)
			}
		}()
	}

	return readLines(input)
}
