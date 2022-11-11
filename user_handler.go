package main

import (
	"final/models"
	"fmt"
	"net/http"
)

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	app.db.Find(&users)
	app.writeJSON(w, http.StatusOK, users, "users")
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var user models.User
	app.db.First(&user, id)
	app.writeJSON(w, http.StatusOK, user, "user")
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	fmt.Println(r)
	app.db.Create(&user)
	app.writeJSON(w, http.StatusOK, user, "user")
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var user models.User
	app.db.First(&user, id)
	app.db.Save(&user)
	app.writeJSON(w, http.StatusOK, user, "user")
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var user models.User
	app.db.First(&user, id)
	app.db.Delete(&user)
	app.writeJSON(w, http.StatusOK, user, "user")
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	app.db.First(&user)
	app.writeJSON(w, http.StatusOK, user, "user")
}
