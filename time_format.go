package gox

import (
	"strings"
)

const (
	stdZeroMonth      = "01"
	stdZeroDay        = "02"
	stdHour           = "15"
	stdZeroMinute     = "04"
	stdZeroSecond     = "05"
	stdLongYear       = "2006"
)

func TimeFormat(format string) string {
	layout := strings.NewReplacer(
		"%YYYY", stdLongYear,
		"%MM", stdZeroMonth,
		"%DD", stdZeroDay,
		"%hh", stdHour,
		"%mm", stdZeroMinute,
		"%ss", stdZeroSecond,
	).Replace(format)

	return layout
}
