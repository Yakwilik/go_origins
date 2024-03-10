package main

import (
	"log"

	"GolangCourse/uniq/handlers"
	"GolangCourse/uniq/io"
	"GolangCourse/uniq/options"
)

func main() {
	opts, err := options.GetOptions()
	if err != nil {
		log.Fatalln(err)
	}

	lines, err := io.GetLines(opts.InputFile)
	if err != nil {
		log.Fatalln(err)
	}

	lines = handlers.HandleLines(lines, opts)
	
	err = io.OutputLines(lines, opts.OutputFile)
	if err != nil {
		log.Fatalln(err)
	}
}
