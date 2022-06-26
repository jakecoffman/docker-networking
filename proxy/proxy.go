package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

const target = "http://example.com"

func main() {
	go handleSignals()

	u, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(u)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Host != "example.com" {
			w.WriteHeader(http.StatusTeapot)
			w.Write([]byte("The proxy only serves http://example.com\n"))
			return
		}
		proxy.ServeHTTP(w, r)
	})

	if err := http.ListenAndServe("0.0.0.0:1080", mux); err != nil {
		log.Fatalln(err)
	}
}

func handleSignals() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM)
	<-ch
	os.Exit(0)
}
