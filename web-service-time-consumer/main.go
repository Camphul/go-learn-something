package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const apiVersion = "0.10.0"
const fetchTimeout = time.Millisecond * 100
const requestTimout = time.Duration(time.Millisecond * 600)

var svcId = "time-consumer-svc"

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
	http.Handle("/", withTimeout(handler)) // each request calls handler
	log.Fatal(http.ListenAndServe(listenerStringFmt, nil))
}
func withTimeout(hand http.HandlerFunc) http.Handler {
	return http.TimeoutHandler(http.HandlerFunc(hand), requestTimout, "Request took too long to complete.")
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - Handling request from %s\n", svcId, r.Host)
	fmt.Printf("%s - Fetching URL %s\n", svcId, timeSvcURL)

	response, err := fetch(r.Context(), timeSvcURL)
	responseVal := response
	if err != nil {

		responseVal = TimestampResponse[string]{true, svcId, fmt.Sprintf("%s", err), apiVersion}
	} else {
		responseVal = TimestampResponse[any]{false, svcId, response, apiVersion}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseVal)
}

func fetch(ctx context.Context, url string) (response any, err error) {
	timeoutContext, cancel := context.WithTimeout(ctx, fetchTimeout)
	defer cancel()
	start := time.Now()
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, url, nil)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}
	defer resp.Body.Close() // don't leak resources
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("while reading %s: %v", url, err)
		return nil, err
	}
	timeSvcResponse := TimestampResponse[int64]{}
	err = json.Unmarshal(b, &timeSvcResponse)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %s", secs, url)
	return timeSvcResponse, nil
}
