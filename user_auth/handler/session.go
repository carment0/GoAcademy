package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSessionCreateHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// when user login they use the POST request, you need to decode to parse the body to the json format to parse it into the struct
		decoder := json.NewDecoder(r.Body)

		var loginReq LoginRequest
		if err := decoder.Decode(&loginReq); err != nil {
			RenderError(w, "Fail to parse request json into a struct", http.StatusInternalServerError)
			return
		}

		//	find user by credentials
		user, err := FindUserByCredential(db, loginReq.Email, loginReq.Password)
		if err != nil {
			RenderError(w, "Incorrect email/password combination", http.StatusUnauthorized)
			return
		}

		expiration := time.Now().Add(2 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session_token", Value: user.SessionToken, Expires: expiration}
		http.SetCookie(w, &cookie)

		res := &UserJSONResponse{
			Name:         user.Name,
			Email:        user.Email,
			SessionToken: user.SessionToken,
		}

		if bytes, err := json.Marshal(res); err != nil {
			// RenderError takes in a string error, since err is a object we have to use .Error to return a string. All error object has this
			RenderError(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}

	}
}

type LogoutResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	isLoggedOut bool   `json:"is_logged_out"`
}

func NewSessionDestroyHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// every time the user submit a http request, the request will have the cookie
		cookie, _ := r.Cookie("session_token")
		if currentUser, err := FindUserByToken(db, cookie.Value); err == nil {
			currentUser.ResetSessionToken()
			db.Save(currentUser)

			res := &LogoutResponse{
				Name:        currentUser.Name,
				Email:       currentUser.Email,
				isLoggedOut: true,
			}

			if bytes, err := json.Marshal(res); err != nil {
				RenderError(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytes)
			}
		} else {
			RenderError(w, "User is not found", http.StatusBadRequest)
		}
	}
}

func NewTokenAuthenticationHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("session_token")
		if currentUser, err := FindUserByToken(db, cookie.Value); err == nil {
			res := UserJSONResponse{
				Name:        currentUser.Name,
				Email:       currentUser.Email,
				SessionToken: currentUser.SessionToken,
			}

			if bytes, err := json.Marshal(res); err != nil {
				RenderError(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytes)
			}
		}  else {
			RenderError(w, err.Error(), http.StatusUnauthorized)
		}
	}
}
