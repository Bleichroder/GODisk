package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/**********************************************************************************
	Database initialization. Creating an connection pool.
**********************************************************************************/
func dbInit() (dbNew *sql.DB, err error) {

	dbNew, err = sql.Open("mysql", "root:/lanxiniu")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Database MySQL connection pool creation succeeded!")
	return dbNew, nil
}

/**********************************************************************************
	Ping MariaDB
**********************************************************************************/
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
	Authority code query
**********************************************************************************/
func authcodeQuery(db *sql.DB, ac string) int {

	var acRet string
	acQuery, err := db.Prepare("SELECT auth_code FROM system_status WHERE disk_name = ?")
	if err != nil {
		log.Println(err)
		return 3
	}
	defer acQuery.Close()
	err = acQuery.QueryRow("demo").Scan(&acRet)
	if err != nil {
		log.Println(err)
		return 3
	}
	if ac != acRet {
		log.Println("The authcode {" + ac + "} does not match with the query value.")
		return 1
	}

	log.Println("The authcode {" + ac + "} is right.")
	return 0
}

/**********************************************************************************
	Username query
**********************************************************************************/
func usernameQuery(db *sql.DB, username string) int {

	var unQueryRet int
	unQuery, err := db.Prepare("SELECT COUNT(*) FROM user_information WHERE username = ?")
	if err != nil {
		log.Println(err)
		return 3
	}
	defer unQuery.Close()
	err = unQuery.QueryRow(username).Scan(&unQueryRet)
	if err != nil {
		log.Println(err)
		return 3
	}
	log.Println("The number of users with the name of {" + username + "} is " + fmt.Sprintf("%d", unQueryRet))
	if unQueryRet != 0 {
		log.Println("User {" + username + "} already exists.")
		return 2
	}

	return 0
}

/**********************************************************************************
	Insert userRegistryInfo into user_information
**********************************************************************************/
func userInfoInsert(db *sql.DB, user *userRegistryInfo) int {

	userInsert, err := db.Prepare("INSERT INTO user_information (username, password, regtime) VALUE(?, ? , now())")
	if err != nil {
		log.Println(err)
		return 3
	}
	defer userInsert.Close()
	_, err = userInsert.Exec(user.name, user.password)
	if err != nil {
		log.Println(err)
		return 3
	}
	log.Println("User information has successfully insert into database.")
	return 0
}

/**********************************************************************************
	Create table user_user.name
**********************************************************************************/
func userTableCreate(db *sql.DB, username string) error {

	var table_name string
	table_name = "user_" + username
	log.Println("Table {" + table_name + "} is creating...")

	createTable, err := db.Prepare("CREATE TABLE " + table_name +
		" (path TEXT(65535) NOT NULL," +
		" file_name TEXT(255) NOT NULL," +
		" type CHAR(20) NOT NULL," +
		" size BIGINT NOT NULL," +
		" modification_time DATETIME NOT NULL," +
		" md5 CHAR(32) NOT NULL," +
		" storage_id BIGINT NOT NULL," +
		" PRIMARY KEY (md5)) ENGINE=Aria")

	if err != nil {
		log.Println(err)
		delRes := userInfoDelete(db, username)
		if delRes != nil {
			log.Println(delRes)
			return delRes
		}
		return err
	}
	defer createTable.Close()
	_, err = createTable.Exec()
	if err != nil {
		log.Println(err)
		return err
	}

	rootfile := &Inode{"/", ".", "folder", 0, time.Now().String(), "md5-1", "1"}
	readme := &Inode{"/", "Readme", "text", 1, time.Now().String(), "md5-2", "2"}
	inodeInsertion(db, table_name, rootfile)
	inodeInsertion(db, table_name, readme)

	log.Println("Table {" + table_name + "} has successfully created.")
	return nil
}

/**********************************************************************************
	User deletion from table user_information
**********************************************************************************/
func userInfoDelete(db *sql.DB, user string) error {

	deletion, err := db.Prepare("DELETE FROM user_information WHERE USERNAME = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	defer deletion.Close()
	_, err = deletion.Exec(user)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("User {" + user + "} has deleted from TABLE user_information.")
	return nil
}

/**********************************************************************************
	Insert file information into table user_username
**********************************************************************************/
func inodeInsertion(db *sql.DB, table string, file *Inode) error {

	log.Println("File infomation of {" + file.FileName + "} inserting...")
	insertQuery, err := db.Prepare("INSERT INTO " + table +
		" (path, file_name, type, size, md5, storage_id, modification_time)" +
		" VALUE(?, ?, ?, ?, ?, ?, now())")

	if err != nil {
		log.Println(err)
		return err
	}
	defer insertQuery.Close()
	_, err = insertQuery.Exec(file.Path, file.FileName, file.FileType, file.FileSize, file.MD5, file.StorageID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("File infomation of {" + file.FileName + "} insertion complete!")
	return nil
}

/**********************************************************************************
	Query file information from table user_username
**********************************************************************************/
func inodeQuery(db *sql.DB, table string, path string) []Inode {

	log.Println("File information query: {table:" + table + "}{path:" + path + "}")

	// Confirm the amount of the files under the path
	var queryRowsLength int
	query, err := db.Prepare("SELECT COUNT(*) From " + table + " WHERE path = \"" + path + "\"")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer query.Close()
	err = query.QueryRow().Scan(&queryRowsLength)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("The number of files under path {" + path + "} is: " + fmt.Sprintf("%d", queryRowsLength))
	result := make([]Inode, queryRowsLength)
	queryRowsCounter := 0

	// Query all files' information from the database
	queryRows, err := db.Query("SELECT * From " + table + " WHERE path = \"" + path + "\"")
	if err != nil {
		log.Println(err)
		return nil
	}
	column, err := queryRows.Columns()
	log.Println("File information column: ", column)
	if err != nil {
		log.Println(err)
		return nil
	}
	values := make([]sql.RawBytes, len(column))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for queryRows.Next() {
		err = queryRows.Scan(scanArgs...)
		if err != nil {
			log.Println(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			switch i {
			case 0:
				result[queryRowsCounter].Path = value
			case 1:
				result[queryRowsCounter].FileName = value
			case 2:
				result[queryRowsCounter].FileType = value
			case 3:
				result[queryRowsCounter].FileSize, _ = strconv.Atoi(value)
			case 4:
				result[queryRowsCounter].ModTime = value
			case 5:
				result[queryRowsCounter].MD5 = value
			case 6:
				result[queryRowsCounter].StorageID = value
			}
		}
		log.Println("File information fetched: ", result[queryRowsCounter])
		queryRowsCounter++
	}
	return result
}

/**********************************************************************************
	Get the amount of registered users.
**********************************************************************************/
func userAmount(db *sql.DB) int {

	var result int
	query, err := db.Prepare("SELECT COUNT(*) FROM user_information")
	if err != nil {
		log.Println(err)
		return 0
	}
	defer query.Close()
	err = query.QueryRow().Scan(&result)
	if err != nil {
		log.Println(err)
		return 0
	}
	log.Println("The amount of users registered is: ", result)
	return result
}

/**********************************************************************************
	Update the amount of registered users.
**********************************************************************************/
func userAmountUpdate(db *sql.DB, disk string, amount int) error {

	query, err := db.Prepare("UPDATE system_status SET users = ? WHERE disk_name = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	defer query.Close()
	_, err = query.Exec(amount, disk)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("The amount of users registered has been updated.")
	return nil
}

/**********************************************************************************
	Update global system statistical information
**********************************************************************************/
func systemStatusUpdate(db *sql.DB) error {

	users := userAmount(db)
	err := userAmountUpdate(db, "demo", users)
	if err != nil {
		return nil
	}
	log.Println("Global system information has been updated.")
	return nil
}

/**********************************************************************************
	Login Services
**********************************************************************************/
func loginService(db *sql.DB, user *userLoginInfo) int {

	// Username query
	var unQueryRet int
	unQuery, err := db.Prepare("SELECT COUNT(*) FROM user_information WHERE username = ?")
	if err != nil {
		log.Println(err)
		return 3
	}
	defer unQuery.Close()
	err = unQuery.QueryRow(user.name).Scan(&unQueryRet)
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
	pwQuery, err := db.Prepare("SELECT password FROM user_information WHERE username = ?")
	if err != nil {
		log.Println(err)
		return 3
	}
	defer pwQuery.Close()
	err = pwQuery.QueryRow(user.name).Scan(&pwQueryRet)
	if err != nil {
		log.Println(err)
		return 3
	}
	log.Println("The password retrieved is {" + pwQueryRet + "}")

	// Password comparision
	if user.password != pwQueryRet {
		log.Println("The password {" + user.password + "} does not match with the query value.")
		return 2
	}

	// Success
	log.Println(user.name + " login success.")
	return 0
}

/**********************************************************************************
	Register Services
**********************************************************************************/
func registerService(db *sql.DB, user *userRegistryInfo) int {

	ret := authcodeQuery(db, user.authcode)
	if ret != 0 {
		return ret
	}

	ret = usernameQuery(db, user.name)
	if ret != 0 {
		return ret
	}

	ret = userInfoInsert(db, user)
	if ret != 0 {
		return ret
	}

	if userTableCreate(db, user.name) != nil {
		return 3
	}

	if systemStatusUpdate(db) != nil {
		return 3
	}

	// Success
	log.Println("User {" + user.name + "} registry success.")
	return 0
}
