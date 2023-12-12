package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/http2"
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

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", *port),
	}

	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		log.Fatal("ConfigureServer", err)
	}

	if *cert == "" || *key == "" {
		log.Printf("server listening at %v without tls", server.Addr)
		log.Fatal(server.ListenAndServe())
	} else {
		log.Printf("server listening at %v with tls", server.Addr)
		log.Fatal(server.ListenAndServeTLS(*cert, *key))
	}

}
