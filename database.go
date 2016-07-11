package main

import (
	"database/sql"
	"fmt"
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

func loginValidate(db *sql.DB, user *userLoginInfo) int {

	// Username query
	var unQueryRet int
	query, err := db.Prepare("SELECT COUNT(*) FROM user_information WHERE username = ?")
	if err != nil {
		log.Println(err)
		return 3
	}
	err = query.QueryRow(user.name).Scan(&unQueryRet)
	if err != nil {
		log.Println(err)
		return 3
	}
	log.Println("The number of users with the name of {" + user.name + "} is " + fmt.Sprintf("%d", unQueryRet))
	if unQueryRet == 0 {
		log.Println(user.name + " does not exist.")
		return 1
	}

	// Password query
	var pwQueryRet string
	query, err = db.Prepare("SELECT password FROM user_information WHERE username = ?")
	if err != nil {
		log.Println(err)
		return 3
	}
	err = query.QueryRow(user.name).Scan(&pwQueryRet)
	if err != nil {
		log.Println(err)
		return 3
	}
	log.Println("The password retrieved is {" + pwQueryRet + "}")

	// Password comparision
	if user.password != pwQueryRet {
		log.Println(user.password + " does not match the query value.")
		return 2
	}

	log.Println(user.name + " login success.")
	return 0
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
