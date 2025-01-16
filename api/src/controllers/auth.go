package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	userSavedOnDB, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusNotFound, errors.New("user not found"))
		return
	}

	if err := security.Compare(userSavedOnDB.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, errors.New("invalid credentials"))
		return
	}

	token, err := auth.CreateToken(userSavedOnDB.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, token)
}
