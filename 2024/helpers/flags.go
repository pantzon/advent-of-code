package helpers

import "flag"

type Flags struct {
	Path *string
}

func PrepCommonFlags() *Flags {
	flags := &Flags{
		Path: flag.String("path", "", ""),
	}
	flag.Parse()
	return flags
}
