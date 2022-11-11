package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	// CRUD for users
	router.HandlerFunc(http.MethodGet, "/api/v1/users", app.getAllUsers)
	router.HandlerFunc(http.MethodGet, "/api/v1/users/:id", app.getUser)
	router.HandlerFunc(http.MethodPost, "/api/v1/users", app.registerUser)
	router.HandlerFunc(http.MethodPut, "/api/v1/users/:id", app.updateUser)
	router.HandlerFunc(http.MethodDelete, "/api/v1/users/:id", app.deleteUser)
	//router.HandlerFunc(http.MethodPost, "/api/v1/users/login", app.loginUser)
	return router
}
