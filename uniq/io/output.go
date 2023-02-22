package io

import (
	"log"
	"os"
)

func OutputLines(lines []string, inputFilename string) (err error) {
	output := os.Stdout

	if inputFilename != "" {
		output, err = os.Create(inputFilename)
		if err != nil {
			return err
		}
		defer func() {
			err = output.Close()
			if err != nil {
				log.Printf("error occured while closing file: %s", err)
			}
		}()
	}
	for _, str := range lines {
		_, err = output.WriteString(str + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
