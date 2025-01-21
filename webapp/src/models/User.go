package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Followers []User    `json:"followers,omitempty"`
	Following []User    `json:"following,omitempty"`
	Posts     []Post    `json:"posts,omitempty"`
}

func SearchFullUser(userID uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go SearchUserData(userChannel, userID, r)
	go SearchFollowers(followersChannel, userID, r)
	go SearchFollowing(followingChannel, userID, r)
	go SearchPosts(postsChannel, userID, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-userChannel:
			if userLoaded.ID == 0 {
				return User{}, errors.New("error on user channel")
			}

			user = userLoaded
		case followersLoaded := <-followersChannel:
			if followersLoaded == nil {
				return User{}, errors.New("error on followers channel")
			}

			followers = followersLoaded
		case followingLoaded := <-followingChannel:
			if followingLoaded == nil {
				return User{}, errors.New("error on following channel")
			}

			following = followingLoaded
		case postsLoaded := <-postsChannel:
			if postsLoaded == nil {
				return User{}, errors.New("error on posts channel")
			}

			posts = postsLoaded
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

func SearchUserData(channel chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.API_URL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- User{}
		return
	}

	channel <- user
}

func SearchFollowers(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.API_URL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err := json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]User, 0) // or just []User{}
		return
	}

	channel <- followers
}

func SearchFollowing(channel chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.API_URL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err := json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]User, 0) // or just []User{}
		return
	}

	channel <- following
}

func SearchPosts(channel chan<- []Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.API_URL, userID)
	response, err := requests.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
    channel <- make([]Post, 0) // or just []Post{}
    return
  }

	channel <- posts
}
