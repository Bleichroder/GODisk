package main

import (
	"fmt"
	"log"
	"net/http"
)

const PCWebDir = "other/web_template"

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func logInHandler(w http.ResponseWriter, r *http.Request) {

}

func signUpHandler(w http.ResponseWriter, r *http.Request) {

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%q\n", r)
}

func main() {
	//http.Handle("/css/", http.FileServer(http.Dir(PCWebDir)))
	//http.Handle("/img/", http.FileServer(http.Dir(PCWebDir)))
	//http.Handle("/js/", http.FileServer(http.Dir(PCWebDir)))
	//http.Handle("/fonts/", http.FileServer(http.Dir(PCWebDir)))
	//http.Handle("/font-awesome/", http.FileServer(http.Dir(PCWebDir)))

	//http.HandleFunc("/", homeHandler)
	http.Handle("/", http.FileServer(http.Dir(PCWebDir)))
	//http.HandleFunc("/login/", logInHandler)
	//http.HandleFunc("/signup/", signUpHandler)
	http.HandleFunc("/index/", indexHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
