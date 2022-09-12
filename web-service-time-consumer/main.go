package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var svcId = "consumer-svc-id"

const apiVersion = "v1"
const timeSvcURL = "http://localhost:8000/timestamp"

type listenerConfig struct {
	host string
	port int
}
type TimestampResponse[v any] struct {
	Error      bool   `json:"error"`
	ServiceId  string `json:"serviceId"`
	Data       v      `json:"data"`
	ApiVersion string `json:"apiVersion"`
}

func main() {
	svcId = fmt.Sprintf("consumer-svc-id-%d", time.Now().UnixMilli())
	listener := listenerConfig{"localhost", 8008}
	listenerStringFmt := fmt.Sprintf("%s:%d", listener.host, listener.port)
	fmt.Printf("Starting time consumer svc http listener on %s\n", listenerStringFmt)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe(listenerStringFmt, nil))
}
func fetch(url string, ch chan<- TimestampResponse[int64]) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err) // send to channel ch
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		fmt.Printf("while reading %s: %v", url, err)
		return
	}
	timeSvcResponse := TimestampResponse[int64]{}
	err = json.Unmarshal(b, &timeSvcResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- timeSvcResponse
	close(ch)
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %s", secs, url)
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	unixMilli := time.Now().UnixMilli()
	fmt.Printf("%d - Handling request from %s\n", unixMilli, r.RemoteAddr)
	fmt.Printf("%d - Fetching URL %s\n", unixMilli, timeSvcURL)
	ch := make(chan TimestampResponse[int64])
	go fetch(timeSvcURL, ch)
	responseVal := TimestampResponse[TimestampResponse[int64]]{false, svcId, <-ch, apiVersion}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseVal)
}
