package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	config.Load()
	cookies.Configure()

	utils.LoadTemplates()
	r := router.Generate()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.WithFields(logrus.Fields{
		"port": config.API_PORT,
	}).Info("🚀 Server is running")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), r))
}
