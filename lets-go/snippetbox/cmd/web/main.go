package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	//	addr := flag.String("addr", ":4000", "HTTP network address")

	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "hTTP network address")

	flag.Parse()

	f, err1 := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)

	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting Server on %s", cfg.addr)

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	//err := http.ListenAndServe(cfg.addr, mux)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}

type config struct {
	addr string
}
