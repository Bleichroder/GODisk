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
}

func (this *registerController) SubmitAction(w http.ResponseWriter, r *http.Request) {
	type Result struct {
		Ret int
		Log string
	}

	var registryResult *Result

	register_username := r.FormValue("register_username")
	register_password := r.FormValue("register_password")
	register_confirm := r.FormValue("register_confirm")
	register_authcode := r.FormValue("register_authcode")

	log.Println("User registry request: " + register_username + "@" + register_password + ":" + register_confirm + "-" + register_authcode)
	if register_authcode != AUTHCODE {
		registryResult = &Result{1, "Wrong authority code!"}
	} else {
		registryResult = &Result{0, "Registry succeeded!"}
	}

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
}

func (this *loginController) SubmitAction(w http.ResponseWriter, r *http.Request) {

	type Result struct {
		Ret int
		Log string
	}

	var loginResult *Result

	login_username := r.FormValue("login_username")
	login_password := r.FormValue("login_password")
	log.Println("User login request: " + login_username + "@" + login_password)

	if (login_username == "admin") && (login_password == "123") {
		loginResult = &Result{0, "Login succeed!"}
	} else {
		loginResult = &Result{1, "Login failed!"}
	}

	b, err := json.Marshal(loginResult)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("Response message hava sent.")
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
}
