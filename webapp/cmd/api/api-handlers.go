package main

import (
	"errors"
	"net/http"
)

type Credentials struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

func (app *Application) authenticate(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	//read a json payload
	err := app.readJSON(w, r, &creds)

	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"))
		return
	}

	//look up the user by email address
	user, err := app.DB.GetUserByEmail(creds.Username)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	//check password
	valid, err := user.PasswordMatches(creds.Password)
	if err != nil || !valid {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	// generate tokens
	tokenPairs, err := app.generateTokenPair(user)
	if err != nil {
		app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
		return
	}

	//send token to user
	_ = app.writeJSON(w, http.StatusOK, tokenPairs)
}

func (app *Application) refresh(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) allUsers(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) getUser(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) deleteUser(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) insertUser(w http.ResponseWriter, r *http.Request) {

}
