package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	os.Setenv("VERSION", "V1.0")
	// http.HandleFunc("/", index)
	http.HandleFunc("/healthz", healthz)
	// mux := http.NewServeMux()
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "ok")
// }

func healthz(w http.ResponseWriter, r *http.Request) {
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
	io.WriteString(w, "Green")
}
