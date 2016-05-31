package main

import (
	"fmt"
	"log"
	"net/http"
)

const PCWebDir = "web/pc"

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
	http.Handle("/css/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/icon/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/jquery/", http.FileServer(http.Dir(PCWebDir)))

	http.HandleFunc("/", homeHandler)
	//http.HandleFunc("/login/", logInHandler)
	http.Handle("/login/", http.FileServer(http.Dir(PCWebDir+"/html/")))
	//http.HandleFunc("/signup/", signUpHandler)
	http.Handle("/signup/", http.FileServer(http.Dir(PCWebDir+"/html/")))
	http.HandleFunc("/index/", indexHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
