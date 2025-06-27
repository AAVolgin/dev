package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	file, err := os.OpenFile("logs/nvr.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.SetOutput(file)
	} else {
		Log.Info("Could not open log file, using default stderr")
	}
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
