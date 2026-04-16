package core

import "errors"



var (
	//400
	ErrBadRequest = errors.New("httpStatusBadRequest")

	//404
	ErrNotFound = errors.New("httpStatusNotFound")

	//409
	ErrConflict = errors.New("httpStatusConflict")

	//500
	ErrInternalServerError = errors.New("httpInternalServerError")
)