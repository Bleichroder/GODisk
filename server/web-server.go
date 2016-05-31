package main

import (
	"fmt"
	"log"
	"net/http"
)

func logInHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello SB!\n")
}

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%q\n", r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/login/", logInHandler)
	http.HandleFunc("/signup/", signUpHandler)
	http.HandleFunc("/index/", indexHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
