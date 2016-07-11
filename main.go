package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const PCWebDir = "static"

func init() {

	// Open Database Server
	GODiskDB, err := dbInit()
	if err != nil {
		log.Println(err)
	}
	GODiskDB.SetMaxOpenConns(2000)
	GODiskDB.SetMaxIdleConns(1000)

	DBPing(GODiskDB)
}

func main() {

	// Web Server
	startWebServer()
}

func startWebServer() {

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
