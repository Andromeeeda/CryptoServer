package core

import "errors"



var (
	//400
	ErrBadRequest = errors.New("httpStatusBadRequest")

	//409
	ErrConflict = errors.New("httpStatusConflict")

	//500
	ErrInternalServerError = errors.New("httpInternalServerError")
)