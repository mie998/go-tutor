package main

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		validateUsername(r.Form.Get("username"))
		validatePassword(r.Form.Get("password"))
		validateAge(r.Form.Get("age"))
	}
}
func validateUsername(str string) {
	r, _ := regexp2.Compile(`^[a-zA-Z0-9]+`, regexp2.RegexOptions(10))
	if m, e := r.MatchString(str); !m || len(e.Error()) > 0 {
		fmt.Println(e)
	} else {
		fmt.Println("username is accepted!!")
	}
}

func validatePassword(str string) {
	r, _ := regexp2.Compile(`^(?=.*?[a-z])(?=.*?[A-Z])(?=.*?\d)[a-zA-Z\d]{8,100}$`, regexp2.RegexOptions(10))
	if m, e := r.MatchString(str); !m || len(e.Error()) > 0 {
		fmt.Println(e)
	} else {
		fmt.Println("password is accepted!!")
	}
}

func validateAge(str string) {
	ageint, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !(0 < ageint && ageint < 100) {
		fmt.Println("your age doesn't seem to be correct.")
		return
	}
	fmt.Println("age: ", ageint)
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
