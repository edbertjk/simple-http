package main

import (
	"fmt"
	"log"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)

	if r.Method != "POST" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error %v", err)
		return
	}
	m["name"] = r.FormValue("name")
	m["kelas"] = r.FormValue("kelas")
	m["umur"] = r.FormValue("umur")
	fmt.Fprintf(w, "Hello %s\n", m["name"])
	fmt.Fprintf(w, "Kelas: %s\n", m["name"])
	fmt.Fprintf(w, "Umur: %s\n", m["umur"])
}

func main() {
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/post", postHandler)
	fmt.Printf("Connect port 80")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
