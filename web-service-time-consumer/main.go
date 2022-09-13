package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const apiVersion = "0.10.0"

var svcId = "time-svc-consumer-id"

var timeSvcURL string

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
	if val, ok := os.LookupEnv("TIMESVC_URL"); ok {
		timeSvcURL = val
	} else {
		timeSvcURL = "http://localhost:8000/timestamp"
	}
	svcId = fmt.Sprintf("%s-%d", svcId, time.Now().UnixMilli())
	listener := listenerConfig{"0.0.0.0", 8008}
	listenerStringFmt := fmt.Sprintf("%s:%d", listener.host, listener.port)
	fmt.Printf("Starting time consumer svc http listener on %s\n", listenerStringFmt)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe(listenerStringFmt, nil))
}
func fetch(url string, ch chan<- TimestampResponse[int64]) {
	start := time.Now()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Print(err) // send to channel ch
		return
	}
	defer close(ch)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err) // send to channel ch
		return
	}
	defer resp.Body.Close() // don't leak resources
	b, err := ioutil.ReadAll(resp.Body)
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
	response := <-ch
	responseVal := TimestampResponse[TimestampResponse[int64]]{false, svcId, response, apiVersion}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseVal)
}
