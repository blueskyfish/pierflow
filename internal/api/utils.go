package api

import (
	"fmt"
	"os"
	"pierflow/internal/logger"
)

func logFromString(s string) (logger.LogLevel, error) {
	switch s {
	case logger.LogNone.String():
		return logger.LogNone, nil
	case logger.LogDebug.String():
		return logger.LogDebug, nil
	case logger.LogInfo.String():
		return logger.LogInfo, nil
	case logger.LogWarn.String():
		return logger.LogWarn, nil
	case logger.LogError.String():
		return logger.LogError, nil
	default:
		return logger.LogNone, fmt.Errorf("invalid log level: %s", s)
	}
}

func listenForExit(quit chan os.Signal) {
	sig := <-quit
	logger.Infof("Received termination signal, exiting with (%s)...", sig.String())
	os.Exit(0)
}

func checkBasePathAndCreatePathIfNecessary(basePath string) string {
	stat, err := os.Stat(basePath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(basePath, 0755); err != nil {
				logger.Errorf("failed to create base path => %s", err.Error())
			}
		} else {
			logger.Errorf("failed to access base path => %s", err.Error())
		}
	} else if !stat.IsDir() {
		logger.Error("base path is not a directory")
	}
	logger.Debugf("base path is set to '%s'", basePath)
	return basePath
}
