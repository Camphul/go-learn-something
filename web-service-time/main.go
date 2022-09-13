package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const apiVersion = "0.10.0"

var svcId = "timestamp-svc"

const requestTimout = time.Duration(time.Millisecond * 50)

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

func withTimeout(hand http.HandlerFunc) http.Handler {
	return http.TimeoutHandler(http.HandlerFunc(hand), requestTimout, "Request took too long to complete.")
}
func main() {
	svcId = fmt.Sprintf("%s-%d", svcId, time.Now().UnixMilli())
	listener := listenerConfig{"0.0.0.0", 8000}
	listenerStringFmt := fmt.Sprintf("%s:%d", listener.host, listener.port)
	fmt.Printf("Starting http listener on %s\n", listenerStringFmt)
	http.Handle("/timestamp", withTimeout(handler))
	http.Handle("/", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(listenerStringFmt, nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().UnixMilli()
	fmt.Printf("%s - Handling %s timestamp request from %s\n", svcId, r.Method, r.Host)
	responseVal := TimestampResponse[int64]{false, svcId, timestamp, apiVersion}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(responseVal)
	if err != nil {
		fmt.Print(err)
	}
}
