package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	addr = flag.String("addr", "http://localhost:8090", "the address to connect to")
	ca   = flag.String("ca", "", "the ca file")
	cert = flag.String("cert", "", "The cert file")
	key  = flag.String("key", "", "The key file")
)

func main() {
	flag.Parse()
	path := fmt.Sprintf("%s/hello", *addr)

	// load tls certificates
	// clientTLSCert, err := tls.LoadX509KeyPair(CertFilePath, KeyFilePath)
	// if err != nil {
	// 	log.Fatalf("Error loading certificate and key file: %v", err)
	// 	return nil
	// }
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	if *ca != "" {
		// Configure the client to trust TLS server certs issued by a CA.
		certPool, err := x509.SystemCertPool()
		if err != nil {
			panic(err)
		}
		if caCertPEM, err := os.ReadFile(*ca); err != nil {
			panic(err)
		} else if ok := certPool.AppendCertsFromPEM(caCertPEM); !ok {
			panic("invalid cert in CA PEM")
		}
		tlsConfig.InsecureSkipVerify = false
		tlsConfig.RootCAs = certPool
	}

	tr := &http2.Transport{
		TLSClientConfig:    tlsConfig,
		DisableCompression: true,
		AllowHTTP:          true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(path)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
	fmt.Printf(
		"Got response %d: %s %s\n",
		resp.StatusCode, resp.Proto, string(body))
}
