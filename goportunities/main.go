package main

import (
	"goportunities/config"
	"goportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// initialize logger
	logger = config.GetLogger("MAIN")

	// initialize configs
	err := config.InitializeAppConfig()

	if err != nil {
		logger.Errorf("ERROR INITIALIZING CONFIGS: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}
