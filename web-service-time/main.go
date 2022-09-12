package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const apiVersion = "1"

var svcId = "svc-id"

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
	svcId = fmt.Sprintf("svc-id-%d", time.Now().UnixMilli())
	listener := listenerConfig{"localhost", 8000}
	listenerStringFmt := fmt.Sprintf("%s:%d", listener.host, listener.port)
	fmt.Printf("Starting http listener on %s\n", listenerStringFmt)
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/timestamp", timestampHandler)
	log.Fatal(http.ListenAndServe(listenerStringFmt, nil))
}

// handler echoes the Path component of the request URL r.
func timestampHandler(w http.ResponseWriter, r *http.Request) {
	unixMilli := time.Now().UnixMilli()
	fmt.Printf("%d - Handling timestamp request from %s\n", unixMilli, r.RemoteAddr)
	responseVal := TimestampResponse[int64]{false, svcId, unixMilli, apiVersion}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseVal)
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	unixMilli := time.Now().UnixMilli()

	fmt.Printf("%d - Handling 404 request from %s\n", unixMilli, r.RemoteAddr)
	responseMsg := fmt.Sprintf("Does not support URL.Path = %q", r.URL.Path)
	responseVal := TimestampResponse[string]{true, svcId, responseMsg, apiVersion}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseVal)
}
