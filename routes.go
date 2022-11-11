package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	// CRUD for users
	router.HandlerFunc(http.MethodGet, "/api/v1/users/", app.getAllUsers)
	router.HandlerFunc(http.MethodGet, "/api/v1/user/", app.getUser)
	router.HandlerFunc(http.MethodPost, "/api/v1/register/", app.registerUser)
	router.HandlerFunc(http.MethodPut, "/api/v1/user/", app.updateUser)
	router.HandlerFunc(http.MethodDelete, "/api/v1/user/", app.deleteUser)
	// login and logout
	//router.HandlerFunc(http.MethodPost, "/api/v1/login", app.loginUser)
	//router.HandlerFunc(http.MethodPost, "/api/v1/logout", app.logoutUser)
	return router
}
