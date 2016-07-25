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

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	log.Println()
	log.Println("Index request: " + action)
	log.Println("Cookie: " + cookie.String())

	http.SetCookie(w, cookie)

	index := &indexController{}
	controller := reflect.ValueOf(index)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		switch action {
		case "":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "indexAction":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "taskAction":
			method = controller.MethodByName(strings.Title("task") + "Action")
		case "settingAction":
			method = controller.MethodByName(strings.Title("setting") + "Action")
		default:
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

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	log.Println()
	log.Println("Registry request: " + action)

	register := &registerController{}
	controller := reflect.ValueOf(register)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		switch action {
		case "":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "indexAction":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "submitAction":
			method = controller.MethodByName(strings.Title("submit") + "Action")
		default:
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

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	log.Println()
	log.Println("Login request: " + action)

	login := &loginController{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		switch action {
		case "":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "indexAction":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "submitAction":
			method = controller.MethodByName(strings.Title("submit") + "Action")
		default:
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
