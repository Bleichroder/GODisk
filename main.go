package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const PCWebDir = "static"

func main() {

	/******************************************************************************
	 MariaDB Client
	******************************************************************************/

	// Open Server
	mariadb, dbOpenErr := sql.Open("mysql", "jason:buck119br@/godisk")
	if dbOpenErr != nil {
		log.Println(dbOpenErr)
	} else {
		log.Println("Database-mysql successfully OPENED!")
	}
	defer mariadb.Close()

	// Ping Server
	pingErr := mariadb.Ping()
	if pingErr != nil {
		log.Println(pingErr)
	} else {
		log.Println("MariaDB successfully validated DSN data!")
	}

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
