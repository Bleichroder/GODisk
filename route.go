package main

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

/**********************************************************************************
 Index path router
**********************************************************************************/
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("username")
	if err != nil {
		log.Println(err)
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	var action = ""
	switch len(parts) {
	case 1:
		action = "IndexAction"
	case 2:
		action = strings.Title(parts[1]) + "Action"
	case 3:
		action = strings.Title(parts[1]) + strings.Title(parts[2]) + "Action"
	default:
		action = ""
	}

	log.Println()
	log.Println("Index request: " + action)
	log.Println("Cookie: " + cookie.String())

	if cookie == nil {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}
	http.SetCookie(w, cookie)

	index := &indexController{}
	controller := reflect.ValueOf(index)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		if action != "" {
			method = controller.MethodByName(action)
		} else {
			NotFoundAction(w)
		}
	}

	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

/**********************************************************************************
 Register path router
**********************************************************************************/
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	var action = ""
	switch len(parts) {
	case 1:
		action = "IndexAction"
	case 2:
		action = strings.Title(parts[1]) + "Action"
	default:
		action = ""
	}

	log.Println()
	log.Println("Registry request: " + action)

	register := &registerController{}
	controller := reflect.ValueOf(register)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		if action != "" {
			method = controller.MethodByName(action)
		} else {
			NotFoundAction(w)
		}
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

/**********************************************************************************
 Login path router
**********************************************************************************/
func LogInHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")

	var action = ""
	switch len(parts) {
	case 1:
		action = "IndexAction"
	case 2:
		action = strings.Title(parts[1]) + "Action"
	default:
		action = ""
	}

	log.Println()
	log.Println("Login request: " + action)

	login := &loginController{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		if action != "" {
			method = controller.MethodByName(action)
		} else {
			NotFoundAction(w)
		}
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

/**********************************************************************************
 Not Found path router
**********************************************************************************/
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

	log.Println()
	log.Println("Not Found request: ")
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
		return
	}
	NotFoundAction(w)
}
