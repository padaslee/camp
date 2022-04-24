package main

import (
	"fmt"
	"log"
	"math/rand"
	"metrics"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// glog.V(2).Info("Starting http server...")
	os.Setenv("VERSION", "V1.0")
	metrics.Register()
	// Use http.DefaultServerMux
	// http.HandleFunc("/", index)
	// http.HandleFunc("/healthz", healthz)
	// err := http.ListenAndServe(":80", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/delay", randomdelay)
	mux.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//copy header
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	//get verion
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	//loging clientip and retcode
	clientip := strings.Split(r.RemoteAddr, ":")[0]
	log.Printf("Client IP: %s, Return Code: 200", clientip)
	//return green
	// io.WriteString(w, "OK")
	fmt.Fprintf(w, "OK")
}

func healthz(w http.ResponseWriter, r *http.Request) {
	//return green
	// io.WriteString(w, "Green")
	fmt.Fprintf(w, "Green")
}

func randomdelay(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}
