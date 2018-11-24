package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"
	"tls"
)

func CopyResHeader(w, res http.Header){

	//type Header map[string][]string
	for key, slc := range res{
		for _, val := range slc {
			w.Add(key, val)
		}
	}
}

func handle(w http.ResponseWriter, req *http.Request){

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	//take the response header and copy it over to our ResponseWriter's header
	CopyResHeader(w.Header(), res.Header)

	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
}

func main (){
	//http proxy functionality - it does not support https right now
	server := &http.Server{
		Addr: "7854",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handle(w, r)
		}),
		//this line disables HTTP/2
		
		//"If TLSNextProto is not nil, HTTP/2 support is not enabled
		// automatically."
		//TLSNextProto map[string]func(*Server, *tls.Conn, Handler) // Go 1.1
		
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
    }

	log.Fatal(server.ListenAndServe())
}