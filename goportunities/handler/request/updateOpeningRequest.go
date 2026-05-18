package requests

import (
	"errors"
	"goportunities/handler/errorHandler"
	"goportunities/schemas"
)

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (request *UpdateOpeningRequest) Validate() error {
	switch {
	case errorHandler.TruncateString(request.Role) == "":
		return errorHandler.ErrorParamIsRequired("role", "string")
	case errorHandler.TruncateString(request.Company) == "":
		return errorHandler.ErrorParamIsRequired("company", "string")
	case errorHandler.TruncateString(request.Location) == "":
		return errorHandler.ErrorParamIsRequired("location", "string")
	case errorHandler.TruncateString(request.Link) == "":
		return errorHandler.ErrorParamIsRequired("link", "string")
	case request.Remote == nil:
		return errorHandler.ErrorParamIsRequired("remote", "boolean")
	case request.Salary <= 0:
		if request.Salary == 0 {
			logger.Errorf("the %s parameter is cannot be zero", "salary")
			return errors.New("salary cannot be zero")
		}
		return errorHandler.ErrorParamIsRequired("salary", "number")
	default:
		return nil
	}
}

func (request *UpdateOpeningRequest) ToEntity() *schemas.Opening {
	return &schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
}