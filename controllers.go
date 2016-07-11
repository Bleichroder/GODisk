package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

/**********************************************************************************
 Index path controller
**********************************************************************************/
type indexController struct{}

func (this *indexController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/500.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Index HTML transmition complete!")
}

func (this *indexController) ErrorAction(w http.ResponseWriter, r *http.Request) {
	NotFoundAction(w)
}

/**********************************************************************************
 Register path controller
**********************************************************************************/

const AUTHCODE = "ppnn13%dkstfeb.1st"

type registerController struct {
}

func (this *registerController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/register.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Register HTML transmition complete!")
}

func (this *registerController) SubmitAction(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Ret int
		Log string
	}

	var registryResult *Result
	var userInfo userRegistryInfo

	// Parameter Transformation
	userInfo.name = r.FormValue("register_username")
	userInfo.password = r.FormValue("register_password")
	userInfo.confirm = r.FormValue("register_confirm")
	userInfo.authcode = r.FormValue("register_authcode")
	log.Println("User registry request: {" + userInfo.name + "}{" + userInfo.password + "}{" + userInfo.confirm + "}{" + userInfo.authcode + "}")

	// Open Database
	GODiskDB, err := dbInit()
	if err != nil {
		log.Println(err)
	}
	defer GODiskDB.Close()

	// Register service
	ret := registerService(GODiskDB, &userInfo)
	switch ret {
	case 0:
		registryResult = &Result{0, "Registry success."}
	case 1:
		registryResult = &Result{1, "Wrong authority code."}
	case 2:
		registryResult = &Result{2, "Username already exists."}
	default:
		log.Println("GODisk server inner error!")
		registryResult = &Result{3, "Inner Error."}
	}

	// Response
	b, err := json.Marshal(registryResult)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Response message have sent.")
		w.Write(b)
	}
}

func (this *registerController) ErrorAction(w http.ResponseWriter, r *http.Request) {
	NotFoundAction(w)
}

/**********************************************************************************
 Login path controller
**********************************************************************************/
type loginController struct {
}

func (this *loginController) IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/html/login.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("Login HTML transmition complete!")
}

func (this *loginController) SubmitAction(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Ret int
		Log string
	}

	var loginResult *Result
	var userInfo userLoginInfo

	// Parameter transformation
	userInfo.name = r.FormValue("login_username")
	userInfo.password = r.FormValue("login_password")
	log.Println("User login request: {" + userInfo.name + "}{" + userInfo.password + "}")

	// Open Database
	GODiskDB, err := dbInit()
	if err != nil {
		log.Println(err)
	}
	defer GODiskDB.Close()

	// Login service
	ret := loginService(GODiskDB, &userInfo)
	switch ret {
	case 0:
		loginResult = &Result{0, "Login success."}
	case 1:
		loginResult = &Result{1, "Username not exist."}
	case 2:
		loginResult = &Result{2, "Password not match."}
	default:
		log.Println("GODisk server inner error!")
		loginResult = &Result{3, "Inner error."}
	}

	// Response
	b, err := json.Marshal(loginResult)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Response message of login handler hava sent.")
		w.Write(b)
	}
}

func (this *loginController) ErrorAction(w http.ResponseWriter, r *http.Request) {
	NotFoundAction(w)
}

/**********************************************************************************
 Not Found path controller
**********************************************************************************/
func NotFoundAction(w http.ResponseWriter) {
	t, err := template.ParseFiles("static/html/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
	log.Println("404 HTML transmition complete!")
}
