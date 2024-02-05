package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

//func snippetCreate(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Create a new snippet..."))
//}

//func snippetCreate(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method != "POST" {
//
//		w.Header().Set("Allow", "POST")
//		w.WriteHeader(405)
//		w.Write([]byte("Method Not Allowed"))
//		return
//
//	}
//
//	w.Write([]byte("Create a new snippet..."))
//}

//func snippetCreate(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method != "POST" {
//
//		w.Header().Set("Allow", "POST")
//		http.Error(w, "Method Not Allowed", 405)
//		return
//
//	}
//
//	w.Write([]byte("Create a new snippet..."))
//}

//func snippetCreate(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method != "POST" {
//
//		w.Header().Set("Allow", http.MethodPost)
//		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		return
//
//	}
//
//	w.Write([]byte("Create a new snippet..."))
//}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.Header()["Date"] = nil
	w.Write([]byte(`{"name":"Alex"}`))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting Server")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
