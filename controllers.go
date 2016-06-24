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

type Result struct {
	Ret int
	Log string
}

func (this *loginController) SubmitAction(w http.ResponseWriter, r *http.Request) {

	var loginSuccess *Result

	login_username := r.FormValue("login_username")
	login_password := r.FormValue("login_password")
	log.Println("User login request: " + login_username + "@" + login_password)

	if (login_username == "admin") && (login_password == "123") {
		loginSuccess = &Result{1, "Login succeed!"}
	} else {
		loginSuccess = &Result{0, "Login failed!"}
	}

	b, err := json.Marshal(loginSuccess)
	if err != nil {
		log.Println(err)
		return
	} else {
		w.Write(b)
		log.Println("Response message hava sent.")
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
