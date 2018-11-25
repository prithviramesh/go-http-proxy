package main

import(
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

)

func main() {
	
	lnk, err := url.Parse("http://localhost:7854")
	
	if err != nil {
		panic(err)
	}

	tr := &http.Transport{
		Proxy : http.ProxyURL(lnk),
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", dump)
}