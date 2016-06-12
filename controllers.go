package main

import (
	"html/template"
	"log"
	"net/http"
)

// Register path controller
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

// Login path controller
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

}

func (this *loginController) ErrorAction(w http.ResponseWriter, r *http.Request) {
	NotFoundAction(w)
}

func NotFoundAction(w http.ResponseWriter) {
	t, err := template.ParseFiles("static/html/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
