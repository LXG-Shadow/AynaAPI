package common

import (
	"AynaAPI/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func createLogger() error {
	Logger = logrus.New()
	file, err := os.OpenFile(config.ServerConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Logger.Out = io.MultiWriter(file, os.Stdout)
	} else {
		Logger.Info("Failed to log to file, using default stdout")
	}
	Logger.SetFormatter(&nested.Formatter{
		HideKeys: false,
		NoColors: true,
	})
	return nil
}
