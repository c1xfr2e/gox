package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_server(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

var (
	keys = []string{"name", "value", "word"}
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	var req inboundRequest
	if err := json.Unmarshal(b, &req); err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	fmt.Printf("%+v\n", req)

	w.WriteHeader(http.StatusOK)
	// query := r.URL.Query()
	// log.Printf("query string: %+v\n", query)
	// for _, k := range keys {
	// 	log.Printf("%s: %s\n", k, query.Get(k))
	// }
	// w.Write([]byte("[current url] : " + currentURL(r) + "\r\n"))
}

func currentURL(r *http.Request) string {
	return r.URL.Path
}

type inboundRequest struct {
	From         string       `json:"From"`
	To           string       `json:"To"`
	CC           string       `json:"Cc"`
	BCC          string       `json:"Bcc"`
	ReplyTo      string       `json:"ReplyTo"`
	Subject      string       `json:"Subject"`
	MessageID    string       `json:"MessageID"`
	ReceivedTime receivedTime `json:"Date"`
	TextBody     string       `json:"TextBody"`
	HTMLBody     string       `json:"HtmlBody"`
	Tag          string       `json:"Tag"`
	Headers      []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Headers"`
	Attachments []struct {
		Name          string `json:"Name"`
		Content       string `json:"Content"`
		ContentType   string `json:"ContentType"`
		ContentLength int    `json:"ContentLength"`
		ContentID     string `json:"ContentID"`
	} `json:"Attachments"`
}

// receivedTime is used to parse `Date` field in request.
type receivedTime time.Time

const timeFormat = "Mon, 1 Jan 2006 15:04:05 -07:00"

var beijing = mustLoadLocation("Asia/Shanghai")

// UnmarshalJSON implements json marshaler interface for Time.
func (t *receivedTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	ti, err := time.ParseInLocation(timeFormat, s, beijing)
	if err != nil {
		return err
	}
	*t = receivedTime(ti)
	return nil
}

func mustLoadLocation(name string) *time.Location {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}
