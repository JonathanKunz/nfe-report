package domain

import "errors"

var (
	InternalServerError = errors.New("Internal Server Error")
	NotFound            = errors.New("Your requested Item is not found")
	Conflict            = errors.New("Your Item already exist")
	BadParamInput       = errors.New("Given Param is not valid")
)
