package main

import (
	"github.com/gorilla/mux" // routing service, multiplexing
	"github.com/jinzhu/gorm"
	"go-academy/user_auth/handler"
	"net/http"
)

func LoadRoutes(db *gorm.DB) http.Handler {
	muxRouter := mux.NewRouter().StrictSlash(true)
	//Name spacing api
	api := muxRouter.PathPrefix("/api").Subrouter()
	api.Handle("/users/authenticate", handler.NewTokenAuthenticationHandler(db)).Methods("GET")
	api.Handle("/users/register", handler.NewUserCreateHandler(db)).Methods("POST")
	api.Handle("/users/login", handler.NewSessionCreateHandler(db)).Methods("POST")
	api.Handle("/users/logout", handler.NewSessionDestroyHandler(db)).Methods("DELETE")
	api.Handle("/users", handler.NewUserListHandler(db)).Methods("GET")
	muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("public"))) // serve static files

	return muxRouter
}
