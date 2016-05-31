package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const PCWebDir = "other/web_template"

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func logInHandler(w http.ResponseWriter, r *http.Request) {

}

func signUpHandler(w http.ResponseWriter, r *http.Request) {

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ls", "-sail", "/home/jason/Documents")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	cmd.Run()
	fmt.Fprintf(w, "%s\n", buffer.String())
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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
