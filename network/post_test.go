package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func makeRequest() {

	message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"name": "value",
		},
	}

	bt, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://webhook.site/#!/", "application/json", bytes.NewBuffer(bt))
	if err != nil {
		log.Fatalln(err)
	}

	strBody, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(strBody))

	//var result map[string]interface{}
	//json.NewDecoder(resp.Body).Decode(&result)
	//log.Println(result["data"])
}

// TestPost does well.
func TestPost(t *testing.T) {
	err := fmt.Errorf("hello: %d", 123)
	fmt.Println(err.Error())
	//makeRequest()
}
