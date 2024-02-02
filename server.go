package main

import (
	"log"
	"net/http"
)

func ZeroTrustServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Never Trust, Always Verify!\n"))
}

func main() {
	http.HandleFunc("/zero", ZeroTrustServer)
	err := http.ListenAndServeTLS(":8443", "certs/acme.example.crt", "certs/acme.example.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
