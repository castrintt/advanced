package errorHandler

import (
	"errors"
	"goportunities/config"
	"strings"
)

var (
	logger = config.NewLogger("ERRORS: ")
)

func ErrorParamIsRequired(name, typeParam string) error {
	logger.Errorf("the %s parameter is required and must be a %s", name, typeParam)
	return errors.New(name + " is required")
}

func TruncateString(text string) string {
	return strings.TrimSpace(text)
}
