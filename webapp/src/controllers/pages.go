package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	utils.RenderTemplate(w, "login.html", nil)
}

func LoadSignupPage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

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

func LoadEditPostPage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.API_URL, postID)
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

	var post models.Post
	if err := json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Err{Err: err.Error()})
		return
	}

	utils.RenderTemplate(w, "update-post.html", post)
}

func LoadUsersPage(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	url := fmt.Sprintf("%s/users?user=%s", config.API_URL, nameOrNick)
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

	var users []models.User
	if err := json.NewDecoder(response.Body).Decode(&users); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.Err{Err: err.Error()})
		return
	}

	utils.RenderTemplate(w, "users.html", users)
}

func LoadUserProfilePage(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Err{Err: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userLoggedID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == userLoggedID {
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	user, err := models.SearchFullUser(userID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Err{Err: err.Error()})
		return
	}

	utils.RenderTemplate(w, "user.html", struct {
		User         models.User
		UserLoggedID uint64
	}{
		User:         user,
		UserLoggedID: userLoggedID,
	})
}

func LoadUserLoggedInProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userLoggedID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, err := models.SearchFullUser(userLoggedID, r)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.Err{Err: err.Error()})
		return
	}

	utils.RenderTemplate(w, "profile.html", user)
}

func LoadEditProfilePage(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Read(r)
	userLoggedID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	channel := make(chan models.User)
	go models.SearchUserData(channel, userLoggedID, r)
	user := <-channel

	if user.ID == 0 {
		responses.JSON(w, http.StatusInternalServerError, responses.Err{Err: "Error on get user"})
		return
	}

	utils.RenderTemplate(w, "update-profile.html", user)
}

func LoadUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {
  utils.RenderTemplate(w, "update-password.html", nil)
}
