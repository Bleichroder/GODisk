package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const PCWebDir = "static"

func main() {

	/******************************************************************************
	 Web Server
	******************************************************************************/
	http.Handle("/css/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/img/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/js/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/fonts/", http.FileServer(http.Dir(PCWebDir)))
	http.Handle("/font-awesome/", http.FileServer(http.Dir(PCWebDir)))

	http.HandleFunc("/index/", IndexHandler)
	http.HandleFunc("/register/", RegisterHandler)
	http.HandleFunc("/login/", LogInHandler)
	http.HandleFunc("/", NotFoundHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
