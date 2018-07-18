package test

import (
	"testing"
	"time"
	"tour/gox"
)

func TestTimeFormat(t* testing.T) {
	sometime := time.Date(2018, 1, 28, 6, 32, 11, 0, time.Local)
	formated := sometime.Format(gox.TimeFormat("%YYYY-%MM-%DD %hh:%mm:%ss"))
	want := "2018-01-28 06:32:11"
	if formated != want {
		t.Errorf("TimeFormat want %s get %s", want, formated)
	}
}
