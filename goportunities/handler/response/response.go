package response

import "goportunities/config"

var (
	logger *config.Logger
)

func InitializeResponse() {
	logger = config.GetLogger("RESPONSE")
	logger.Info("Initializing response...")
}
