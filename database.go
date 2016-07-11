package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var GODiskDB *sql.DB

// Database initialization. Creating an connection pool.
func dbInit() (dbNew *sql.DB, err error) {

	dbNew, err = sql.Open("mysql", "jason:buck119br@/godisk")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Database MySQL connection pool creation succeeded!")
	log.Println(dbNew)
	return dbNew, nil
}

// Ping MariaDB
func DBPing(db *sql.DB) {

	var err error

	if db != nil {
		err = db.Ping()
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		log.Println("NIL connection pool pointer!")
		return
	}

	log.Println("MariaDB successfully validated DSN data!")
}

/**********************************************************************************
	Login Services
**********************************************************************************/

type userLoginInfo struct {
	name     string
	password string
}

func loginValidate(db *sql.DB, user *userLoginInfo) bool {

	var queryRet string
	query, err := db.Prepare("SELECT password FROM user_information WHERE username = ?")
	if err != nil {
		log.Println(err)
		return false
	}

	err = query.QueryRow(user.name).Scan(&queryRet)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("The password retrieved is {" + queryRet + "}")
	return true
}

/**********************************************************************************
	Register Services
**********************************************************************************/

type userRegistryInfo struct {
	name     string
	password string
	confirm  string
	authcode string
}
