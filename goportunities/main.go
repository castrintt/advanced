package main

import (
	"goportunities/config"
	"goportunities/router"
)

func main() {
	// initialize logger
	logger := config.NewLogger("GO-OPPORTUNITIES")


	// initialize configs
	err := config.Initialize()

	if err != nil {
		logger.Error(err)
		return
	}

	// initialize router
	router.Initialize()
}
