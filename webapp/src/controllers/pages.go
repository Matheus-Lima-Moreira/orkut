package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.html", nil)
}

func LoadSignupPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "signup.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.API_URL)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Err{Err: err.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleErrorStatusCode(w, response)
		return
	}

	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Err{Err: err.Error()})
		return
	}

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		fmt.Println("Erro ao carregar fuso horário:", err)
		return
	}

	for i := range posts {
		utcTime := posts[i].CreatedAt.Format("2006-01-02 15:04:05")
		layout := "2006-01-02 15:04:05"

		t, err := time.ParseInLocation(layout, utcTime, time.UTC)
		if err != nil {
			fmt.Println("Erro ao analisar a data:", err)
			return
		}

		posts[i].CreatedAt = t.In(loc)
	}

	cookie, _ := cookies.Read(r)

	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.RenderTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}
