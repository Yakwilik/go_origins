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
	flag.BoolVar(&opts.EShowStrMeetCount, eShowStrMeetCountFlag, opts.EShowStrMeetCount, showStrMeetCountFlagUsage)
	flag.BoolVar(&opts.EShowNotUniqueStr, eShowNotUniqueStrFlag, opts.EShowNotUniqueStr, showNotUniqueStrFlagUsage)
	flag.BoolVar(&opts.EShowUniqueStr, eShowUniqueStrFlag, opts.EShowUniqueStr, showUniqueStrFlagUsage)
	flag.IntVar(&opts.SkippedStringsCount, skippedStringsCountFlag, opts.SkippedStringsCount, skippedStringsCountFlagUsage)
	flag.IntVar(&opts.SkippedCharsCount, skippedCharsCountFlag, opts.SkippedCharsCount, skippedCharsCountFlagUsage)
	flag.BoolVar(&opts.IgnoreRegister, ignoreRegisterFlag, opts.IgnoreRegister, ignoreRegisterUsage)
	flag.Parse()
	opts.InputFile = flag.Arg(0)
	opts.OutputFile = flag.Arg(1)
	return opts, opts.validateOptions()
}

/*
 */
func (o *Options) validateOptions() (err error) {
	exclusiveFlagMet := 0

	checkFlag := func(f bool, count *int) {
		if f {
			*count++
		}
	}
	checkFlag(o.EShowStrMeetCount, &exclusiveFlagMet)
	checkFlag(o.EShowNotUniqueStr, &exclusiveFlagMet)
	checkFlag(o.EShowUniqueStr, &exclusiveFlagMet)
	if exclusiveFlagMet > 1 {
		return fmt.Errorf("only one of these flags are possible at one: %s, %s, %s",
			eShowStrMeetCountFlag, eShowNotUniqueStrFlag, eShowUniqueStrFlag)
	}
	return err
}
