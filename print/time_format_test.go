package print

import (
	"strings"
	"testing"
	"time"
)

const (
	stdZeroMonth  = "01"
	stdZeroDay    = "02"
	stdHour       = "15"
	stdZeroMinute = "04"
	stdZeroSecond = "05"
	stdLongYear   = "2006"
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

func TestTimeFormat(t *testing.T) {
	sometime := time.Date(2018, 1, 28, 6, 32, 11, 0, time.Local)
	formated := sometime.Format(TimeFormat("%YYYY-%MM-%DD %hh:%mm:%ss"))
	want := "2018-01-28 06:32:11"
	if formated != want {
		t.Errorf("TimeFormat want %s get %s", want, formated)
	}
}
