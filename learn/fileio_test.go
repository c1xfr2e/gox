package learn

import (
	"fmt"
	"os"
	"testing"
)

func TestFileIO(t *testing.T) {
	fmt.Println("start")
	f, err := os.Open("data/people.pb")
	if err != nil {
		fmt.Println(err)
		return
	}
	fi, err := f.Stat()
	b := make([]byte, fi.Size())
	_, err = f.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
