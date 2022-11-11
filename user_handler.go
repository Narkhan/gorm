package main

import (
	"final/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	app.db.Find(&users)
	app.writeJSON(w, http.StatusOK, users, "users")
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	var user models.User
	app.db.First(&user, id)
	app.writeJSON(w, http.StatusOK, user, "user")
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	r.ParseForm()
	user.Name = r.Form.Get("name")
	user.Email = r.Form.Get("email")
	user.Phone = r.Form.Get("phone")
	user.Password = r.Form.Get("password")
	// crypt password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		app.errorLog.Println(err)
	}
	user.Password = string(hashPassword)

	app.db.Create(&user)
	app.writeJSON(w, http.StatusOK, user, "user")
}

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	var user models.User
	app.db.First(&user, id)
	r.ParseForm()

	if r.Form.Get("name") != "" {
		user.Name = r.Form.Get("name")
	}
	if r.Form.Get("email") != "" {
		user.Email = r.Form.Get("email")
	}
	if r.Form.Get("phone") != "" {
		user.Phone = r.Form.Get("phone")
	}
	if r.Form.Get("password") != "" {
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.Form.Get("password")), bcrypt.DefaultCost)
		if err != nil {
			app.errorLog.Println(err)
		}
		user.Password = string(hashPassword)
	}

	app.db.Save(&user)
	app.writeJSON(w, http.StatusOK, user, "user")
}

func (app *application) deleteUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	var user models.User
	app.db.First(&user, id)
	// check if user exists or not equal to id
	if user.ID == 0 {
		app.writeJSON(w, http.StatusNotFound, nil, "user")
		return
	}
	app.db.Delete(&user)
	app.writeJSON(w, http.StatusOK, user, "user")
}

// Login and logout
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	app.db.Where("email = ?", r.Form.Get("email")).First(&user)
	if user.ID == 0 {
		app.writeJSON(w, http.StatusNotFound, nil, "user")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Form.Get("password")))
	if err != nil {
		app.writeJSON(w, http.StatusUnauthorized, nil, "user")
		return
	}
	app.writeJSON(w, http.StatusOK, user, "user")
}
