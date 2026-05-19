package requests

import (
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
	case request.Role != "" && errorHandler.TruncateString(request.Role) == "":
		return errorHandler.ErrorParamIsRequired("role", "string")
	case request.Company != "" && errorHandler.TruncateString(request.Company) == "":
		return errorHandler.ErrorParamIsRequired("company", "string")
	case request.Location != "" && errorHandler.TruncateString(request.Location) == "":
		return errorHandler.ErrorParamIsRequired("location", "string")
	case request.Link != "" && errorHandler.TruncateString(request.Link) == "":
		return errorHandler.ErrorParamIsRequired("link", "string")
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
