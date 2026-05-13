package requests

import (
	"errors"
	"strings"
)

func errorParamIsRequired(name, typeParam string) error {
	logger.Errorf("the %s parameter is required and must be a %s", name, typeParam)
	return errors.New(name + " is required")
}

func truncateString(str string) string {
	return strings.TrimSpace(str)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (request *CreateOpeningRequest) Validate() error {
	switch {
	case truncateString(request.Role) == "":
		return errorParamIsRequired("role", "string")
	case truncateString(request.Company) == "":
		return errorParamIsRequired("company", "string")
	case truncateString(request.Location) == "":
		return errorParamIsRequired("location", "string")
	case truncateString(request.Link) == "":
		return errorParamIsRequired("link", "string")
	default:
		return nil
	}
}
