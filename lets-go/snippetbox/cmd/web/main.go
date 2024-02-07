package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	//	addr := flag.String("addr", ":4000", "HTTP network address")

	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "hTTP network address")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting Server on %s", cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)

}

type config struct {
	addr string
}
