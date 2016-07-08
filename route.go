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

	log.Println()
	log.Print("Index request:")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	index := &indexController{}
	controller := reflect.ValueOf(index)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		switch action {
		case "":
			method = controller.MethodByName(strings.Title("index") + "Action")
		default:
			method = controller.MethodByName(strings.Title("error") + "Action")
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

	log.Println()
	log.Println("Registry request:")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	register := &registerController{}
	controller := reflect.ValueOf(register)
	method := controller.MethodByName(action)

	if !method.IsValid() {
		switch action {
		case "":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "index":
			method = controller.MethodByName(strings.Title("index") + "Action")
		case "submit":
			method = controller.MethodByName(strings.Title("submit") + "Action")
		default:
			method = controller.MethodByName(strings.Title("error") + "Action")
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

	log.Println()
	log.Println("Login Registry:")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")

	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

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
			method = controller.MethodByName(strings.Title("error") + "Action")
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

	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}
	NotFoundAction(w)
}
