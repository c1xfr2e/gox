package jsonn

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const msg = `{"status":"OK","errMsg":null,"errCode":0,"data":{"courier":"18401464513","orderNo":"DH2018091510494577","trialMessage":"[{\"time\":\"2018-09-15 16:57:06\",\"lng\":116.484618,\"lat\":39.997286},{\"time\":\"2018-09-15 16:59:23\",\"lng\":116.485153,\"lat\":39.997796},{\"time\":\"2018-09-15 16:59:38\",\"lng\":116.485851,\"lat\":39.998522},{\"time\":\"2018-09-15 16:59:53\",\"lng\":116.487202,\"lat\":39.999736},{\"time\":\"2018-09-15 17:00:08\",\"lng\":116.488333,\"lat\":40.000735},{\"time\":\"2018-09-15 17:00:23\",\"lng\":116.489967,\"lat\":40.001104},{\"time\":\"2018-09-15 17:00:39\",\"lng\":116.491266,\"lat\":40.0012},{\"time\":\"2018-09-15 17:00:54\",\"lng\":116.493168,\"lat\":40.001408},{\"time\":\"2018-09-15 17:01:09\",\"lng\":116.495044,\"lat\":40.001633},{\"time\":\"2018-09-15 17:01:24\",\"lng\":116.495002,\"lat\":40.002935},{\"time\":\"2018-09-15 17:01:39\",\"lng\":116.494669,\"lat\":40.003928},{\"time\":\"2018-09-15 17:05:58\",\"lng\":116.495164,\"lat\":40.002125},{\"time\":\"2018-09-15 17:06:13\",\"lng\":116.493471,\"lat\":40.001681},{\"time\":\"2018-09-15 17:06:28\",\"lng\":116.492014,\"lat\":40.001522},{\"time\":\"2018-09-15 17:06:44\",\"lng\":116.490479,\"lat\":40.00126},{\"time\":\"2018-09-15 17:06:59\",\"lng\":116.489931,\"lat\":39.999658},{\"time\":\"2018-09-15 17:07:14\",\"lng\":116.489032,\"lat\":39.998547},{\"time\":\"2018-09-15 17:07:29\",\"lng\":116.487838,\"lat\":39.997308},{\"time\":\"2018-09-15 17:07:44\",\"lng\":116.486257,\"lat\":39.99572},{\"time\":\"2018-09-15 17:08:00\",\"lng\":116.484622,\"lat\":39.994222},{\"time\":\"2018-09-15 17:09:15\",\"lng\":116.483895,\"lat\":39.993645},{\"time\":\"2018-09-15 17:09:31\",\"lng\":116.483374,\"lat\":39.99324},{\"time\":\"2018-09-15 17:09:46\",\"lng\":116.48271,\"lat\":39.992772},{\"time\":\"2018-09-15 17:10:01\",\"lng\":116.482085,\"lat\":39.992342},{\"time\":\"2018-09-15 17:10:16\",\"lng\":116.481499,\"lat\":39.991873},{\"time\":\"2018-09-15 17:10:31\",\"lng\":116.480893,\"lat\":39.991409},{\"time\":\"2018-09-15 17:10:47\",\"lng\":116.48029,\"lat\":39.990886},{\"time\":\"2018-09-15 17:12:17\",\"lng\":116.479835,\"lat\":39.990561}]"}}`

// Address defines address struct.
type Address struct {
	City string  `json:"city"` // 城市
	Addr string  `json:"addr"` // 地址(街道、小区、大厦等，用于定位)
	Lng  float64 `json:"lng"`
	Lat  float64 `json:"lat"`
}

// Person defines person struct.
type Person struct {
	Name    string     `json:"name"`
	Number  int32      `json:"number"`
	Address []*Address `json:"addressList"`
}

type sPerson struct {
	Person
	Signature string `json:"signature"`
}

func struct2map(st interface{}) (map[string]interface{}, error) {
	b, err := json.Marshal(&st)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func mapFromStruct(s interface{}) map[string]string {
	m := make(map[string]string)
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		jt, ok := val.Type().Field(i).Tag.Lookup("json")
		if !ok {
			continue
		}
		m[jt] = val.Field(i).String()
	}
	return m
}

// S2J try converting struct to json.
func S2J() string {
	p := Person{
		Name:   "Sprocket",
		Number: 2121,
		Address: []*Address{
			&Address{
				City: "北京市",
				Addr: "北京市东城区张自忠路",
				Lng:  116.313545,
				Lat:  40.042105,
			},
			&Address{
				City: "北京市",
				Addr: "北京市朝阳区望京浦项中心7层AB",
			},
		},
	}
	e, err := json.Marshal(&p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(e))

	sp := sPerson{
		Person:    p,
		Signature: "A2A13BVKD-LS9J$F#EI3OTAL",
	}
	d, err := json.Marshal(&sp)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	m, err := struct2map(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(m)

	return string(d)
}
