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

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		/* w.Header().Set("Cache-Control", "public, max-age=31536000")
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")
		w.Header().Del("Cache-Control")
		w.Header().Get("Cache-Control") */
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/snippet", showSnippet)
	http.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}
