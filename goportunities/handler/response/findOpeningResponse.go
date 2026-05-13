package response

import (
	"goportunities/schemas"
)

type OpeningResponse struct {
	ID       uint   `json:"id"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (response *OpeningResponse) FromEntity(entity *schemas.Opening) *OpeningResponse {
	response.ID = entity.ID
	response.Role = entity.Role
	response.Company = entity.Company
	response.Location = entity.Location
	response.Remote = &entity.Remote
	response.Link = entity.Link
	response.Salary = entity.Salary
	return response
}
