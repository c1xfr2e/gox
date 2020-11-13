package learn

import "testing"

func TestScope(t *testing.T) {
	str := ""
	if str == "" {
		// This will shadow outer `str` and initialize an `str` in inner scope.
		// str, err := getStr()
		var err error
		str, err = getStr()
		if err != nil {
			t.Error(err)
		}
		t.Logf("str=%q\n", str)
	}
	t.Logf("str=%q\n", str)
}

func getStr() (string, error) {
	return "hello", nil
}
