package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const Addr = ":3000"

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	_, err := SetupDatabase()

	if err != nil {
		logrus.Error(err)
		return
	}

	// create server with http package
	server := &http.Server{
		Addr: Addr,
		Handler: LoadRoutes(),
		// the thread the writes or read the handle will terminate if over 15 sec
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Infof("HTTP server is listening and serving on port %v", Addr)
	logrus.Fatal(server.ListenAndServe())
}
