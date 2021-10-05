package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func readyForXSS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("xss: ", r.Form["ready"])
	}
}
func notReadyForXSS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("xss: ", r.Form["notReady"])
	}
}

func main() {
	http.HandleFunc("/readyForXSS", readyForXSS)
	http.HandleFunc("/notReadyForXSS", notReadyForXSS)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAdnServe: ", err)
	}
}
