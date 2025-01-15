package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	userId, err := repository.Create(user)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = userId

	responses.JSON(w, http.StatusCreated, user)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	users, err := repository.List(nameOrNick)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func ShowUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing User!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User!"))
}
