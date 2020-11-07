package domain

import "errors"

var (
	//ErrNotfound ...
	ErrNotfound = errors.New("Your Request is not found")
	//ErrInternalServerError ...
	ErrInternalServerError = errors.New("Internal Server Error")
	//ErrUnauthorized ...
	ErrUnauthorized = errors.New("Your Request Is unauthorized")
)
