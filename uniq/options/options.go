package options

import (
	"flag"
	"fmt"
)

type Options struct {
	EShowStrMeetCount   bool
	EShowNotUniqueStr   bool
	EShowUniqueStr      bool
	SkippedStringsCount int
	SkippedCharsCount   int
	IgnoreRegister      bool
	InputFile           string
	OutputFile          string
}

func GetOptions() (opts Options, err error) {
	flag.BoolVar(&opts.EShowStrMeetCount, "c", false, showStrMeetCountFlagUsage)
	flag.BoolVar(&opts.EShowNotUniqueStr, "d", false, showNotUniqueStrFlagUsage)
	flag.BoolVar(&opts.EShowUniqueStr, "u", false, showUniqueStrFlagUsage)
	flag.IntVar(&opts.SkippedStringsCount, "f", 0, skippedStringsCountFlagUsage)
	flag.IntVar(&opts.SkippedCharsCount, "s", 0, skippedCharsCountFlagUsage)
	flag.BoolVar(&opts.IgnoreRegister, "i", false, ignoreRegisterUsage)
	flag.Parse()
	opts.InputFile = flag.Arg(0)
	opts.OutputFile = flag.Arg(1)
	return opts, opts.validateOptions()
}

func (o *Options) validateOptions() (err error) {
	exclusiveFlagsCount := 0

	checkFlag := func(isFlagTrue bool) {
		if isFlagTrue {
			exclusiveFlagsCount++
		}
	}
	checkFlag(o.EShowStrMeetCount)
	checkFlag(o.EShowNotUniqueStr)
	checkFlag(o.EShowUniqueStr)
	if exclusiveFlagsCount > 1 {
		return fmt.Errorf("only one of these flags are possible at ones: -c, -d, -u")
	}
	return err
}
