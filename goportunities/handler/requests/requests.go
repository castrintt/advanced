package requests

import "goportunities/config"

var (
	logger *config.Logger
)

func InitializeRequests() {
	logger = config.GetLogger("REQUESTS")
}
