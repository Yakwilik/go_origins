package structs

/* {"errors":[{"location":"body","param":"username","value":"yakwilik","msg":"already exists"}]} */

type ErrorResponse struct {
	Err []ErrorMap `json:"errors"`
}

type ErrorMap map[string]string

func NewErrorResponse() ErrorResponse {
	return ErrorResponse{Err: []ErrorMap{}}
}

func NewErrorMap() ErrorMap {
	return ErrorMap{}
}
