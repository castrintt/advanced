package requests

import (
	"errors"
	"goportunities/schemas"
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
	case request.Remote == nil:
		return errorParamIsRequired("remote", "boolean")
	case request.Salary <= 0:
		return errorParamIsRequired("salary", "number")
	default:
		return nil
	}
}

func (request *CreateOpeningRequest) ToEntity() *schemas.Opening {
	return &schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
}
