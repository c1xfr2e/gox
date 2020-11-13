package learn

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func listStructFields(obj interface{}) {
	val := reflect.ValueOf(obj).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

// TestReflect does of testing of reflect
func TestReflect1(t *testing.T) {
	type User struct {
		Name string `tag_name:"t-name"`
		Age  int    `tag_name:"t-age"`
	}
	listStructFields(&User{Name: "name", Age: 32})
}

type obj struct {
	Key1 string `json:"k1"`
	Key2 string `json:"k2"`
	Key3 int64  `json:"k3"`
	Key4 int    `json:"k4"`
	Key5 bool   `json:"k5"`
}

func TestReflect2(t *testing.T) {
	data := `{"k1": "v1", "k2": "v2", "k3": 1234567890, "k4": 456, "k5": true}`
	d := map[string]interface{}{}
	json.Unmarshal([]byte(data), &d)
	obj := &obj{}
	s := reflect.ValueOf(obj).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		for j, f := range d {
			if typeOfT.Field(i).Tag.Get("json") == j {
				fl := s.FieldByName(typeOfT.Field(i).Name)
				switch fl.Kind() {
				case reflect.Bool:
					fl.SetBool(f.(bool))
				case reflect.Int, reflect.Int64:
					c, _ := f.(float64)
					fl.SetInt(int64(c))
				case reflect.String:
					fl.SetString(f.(string))
				}
			}
		}
	}
	fmt.Printf("%+v\n", obj) // &{Key1:v1 Key2:v2 Key3:1234567890 Key4:456 Key5:true}
}
