package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.Printf("output: hello\n")
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

var (
	port = flag.Int("port", 8090, "The server port")
	cert = flag.String("cert", "", "The cert file")
	key  = flag.String("key", "", "The key file")
)

func main() {
	flag.Parse()
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	addr := fmt.Sprintf(":%d", *port)
	if *cert == "" || *key == "" {
		log.Printf("server listening at %v without tls", addr)
		log.Fatal(http.ListenAndServe(addr, nil))
	} else {
		log.Printf("server listening at %v with tls", addr)
		log.Fatal(http.ListenAndServeTLS(addr, *cert, *key, nil))
	}

}
