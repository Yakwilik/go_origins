package main

import (
	"GolangCourse/uniq/handlers"
	"GolangCourse/uniq/io"
	"GolangCourse/uniq/options"
	"fmt"
	"os"
)

func main() {
	opts, err := options.GetOptions()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	lines, err := io.GetLines(opts.InputFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	lines = handlers.HandleLines(lines, opts)
	for _, str := range lines {
		fmt.Println(str)
	}
}
