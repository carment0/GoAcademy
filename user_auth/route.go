package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func LoadRoutes() http.Handler {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	return muxRouter
}