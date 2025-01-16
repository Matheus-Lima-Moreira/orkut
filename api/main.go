package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	config.Load()

	r := router.Generate()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.WithFields(logrus.Fields{
		"port": config.Port,
	}).Info("🚀 Server is running")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
