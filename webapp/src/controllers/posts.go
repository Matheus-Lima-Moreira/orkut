package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"Title":   r.FormValue("title"),
		"Content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts", config.API_URL)
	response, err := requests.RequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(post))
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

func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/like", config.API_URL, postID)
	response, err := requests.RequestWithAuthentication(r, http.MethodPost, url, nil)
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

func DislikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	postID, err := strconv.ParseUint(parameters["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.Err{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/dislike", config.API_URL, postID)
	response, err := requests.RequestWithAuthentication(r, http.MethodPost, url, nil)
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
