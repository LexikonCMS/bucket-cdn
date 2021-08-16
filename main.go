package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// Get configuration from environment
	hostname, hasHost := os.LookupEnv("HOST")
	bucketName, hasBucket := os.LookupEnv("BUCKET_NAME")
	if !hasBucket {
		log.Fatalln("No bucket name set in environment.")
	}

	// Create remote url
	remote, err := url.Parse(fmt.Sprintf("https://storage.googleapis.com/%s", bucketName))
	if err != nil {
		log.Fatalln(err)
	}

	// Build address (default is :8080)
	var addr string
	if addr = ":8080"; hasHost {
		addr = fmt.Sprintf("%s:8080", hostname)
	}

	// Define reverse handler function
	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Host = remote.Host
			w.Header().Set("X-Powered-By", "Lexikon BucketCDN")
			p.ServeHTTP(w, r)
		}
	}

	// Create proxy and start listening
	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	log.Printf("Serving %s on %s", bucketName, addr)

	if err = http.ListenAndServe(addr, nil); err != nil {
		log.Fatalln(err)
	}
}
