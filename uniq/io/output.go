package io

import (
	"log"
	"os"
)

func OutputLines(lines []string, outputFilename string) (err error) {
	output := os.Stdout

	if outputFilename != "" {
		output, err = os.Create(outputFilename)
		if err != nil {
			return err
		}
		defer func() {
			closeErr := output.Close()
			if err != nil {
				log.Printf("error occured while closing file: %s", closeErr)
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
