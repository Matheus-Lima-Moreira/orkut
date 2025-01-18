package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"nick":     r.FormValue("nick"),
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.API_URL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Err{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/auth/login", config.API_URL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(user))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Err{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var dataAuthentication models.DataAuthentication
	if err := json.NewDecoder(response.Body).Decode(&dataAuthentication); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Err{Err: err.Error()})
	}

	if err := cookies.Save(w, dataAuthentication.ID, dataAuthentication.Token); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Err{Err: err.Error()})
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
