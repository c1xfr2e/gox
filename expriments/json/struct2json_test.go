package jsonn

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func tryKV() (k []string, kv map[string]string) {
	k = append(k, "foo")
	k = append(k, "bar")
	kv = make(map[string]string)
	kv["foo"] = "bar"
	return
}

func TestS2J(t *testing.T) {
	k, kv := tryKV()
	fmt.Println(k)
	fmt.Println(kv)

	s := S2J()
	t.Log(s)

	m := make(map[string]string)
	m["name"] = "value"
	m["hello"] = "world"
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}

type Foo struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

func toString(v interface{}) (string, error) {
	s, ok := v.(string)
	if ok {
		return s, nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func TestMarshal(t *testing.T) {
	foo := Foo{
		Name:  "name",
		Value: 1,
	}
	t.Log(toString("sss"))
	t.Log(toString(123))
	t.Log(toString(foo))

	a := []int{12, 21}
	s, err := json.Marshal(a)
	t.Logf("Failed to marshal %v, err=%v", a, err)
	t.Log(string(s))

	m := map[string]string{
		"sign":       "210833c7af5c03aa205e48cb7e3ed44a",
		"param_json": `{"page":"0","size":"1"}`,
		"method":     "product.list",
		"app_key":    "3311984797114607023",
		"timestamp":  "2019-06-13 14:52:54",
		"v":          "1",
	}
	ms, _ := json.Marshal(m)
	t.Log(string(ms))
}

func TestEmbedded(t *testing.T) {
	var foo Foo
	s := &struct {
		*Foo
		Value string `json:"value"`
		Alias string `json:"alias"`
	}{
		Foo: &foo,
	}
	err := json.Unmarshal([]byte(`{"name":"BigPowerLiu", "value": "22"}`), s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", s.Foo)
	i, err := strconv.Atoi(s.Value)
	if err != nil {
		t.Fatal(err)
	}
	foo.Value = int32(i)
	fmt.Printf("%v\n", foo)
}
