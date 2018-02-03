package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const Addr = ":3000"

func main() {
	// logrus: printing library, print with color in terminal and saves logs
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	db, err := SetupDatabase()

	if err != nil {
		logrus.Error(err)
		return
	}

	// create server with http package
	// we are using a server struct, allow more configuring
	server := &http.Server{
		Addr:    Addr,
		Handler: LoadRoutes(db), // attach routes to server
		// the thread the writes or read the handle, will terminate after 15 sec
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("HTTP server is listening and serving on port %v", Addr)
	logrus.Fatal(server.ListenAndServe())
}
