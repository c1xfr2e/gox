package learn

import (
	"encoding/json"
	"fmt"
	"testing"
)

type S1 struct {
	Value string `json:"value"`
}

type Alias S1

type S2 struct {
	A1    *Alias `json:"foo1"`
	A2    *Alias `json:"foo2"`
	Value string `json:"value"`
}

func (a *Alias) UnmarshalJSON(b []byte) error {
	var f float64
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}
	*a = (Alias)(S1{
		Value: fmt.Sprintf("%f", f),
	})
	return nil
}

func TestUnmarshal(t *testing.T) {
	var s2 S2
	err := json.Unmarshal([]byte(`{"foo1":3.14, "foo2":0.618, "value": "barbarbar"}`), &s2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v\n", s2.A1)
	fmt.Printf("%+v\n", s2.A2)
}
