package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

var MariaDB *MySQL

// Database Initialization
func dbInit() (dbNew *MySQL, err error) {

	dbNew = new(MySQL)
	dbNew.db, err = sql.Open("mysql", "jason:buck119br@/godisk")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Database-mysql successfully OPENED!")
	return dbNew, nil
}

// Ping MariaDB
func (db *MySQL) ping() {

	var err error

	if db != nil {
		err = db.db.Ping()
	} else {
		log.Println("Not available database pointer!")
		return
	}

	if err != nil {
		log.Println(err)
	}

	log.Println("MariaDB successfully validated DSN data!")
}

// Database close
func (db *MySQL) Close() {
	log.Println("Database-mysql will be closed!")
	db.db.Close()
}

/**********************************************************************************
	Login Services
**********************************************************************************/

type userLoginInfo struct {
	name     string
	password string
}

func (db *MySQL) loginValidate(user *userLoginInfo) (result bool) {

	var queryRet string
	query, err := db.db.Prepare("SELECT password FROM user_information WHERE username = ?")
	if err != nil {
		log.Println(err)
		return false
	}

	err = query.QueryRow(user.name).Scan(&queryRet)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("the password retrieved is " + queryRet)
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
