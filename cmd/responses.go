package cmd

import "net/http"

type (
	responseType struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
		Error   *errorWrap  `json:"error"`
	}

	errorWrap struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
	}

	defaultErr struct{}
)

var (
	ForbiddenResponse           = responseType{false, nil, &errorWrap{"Forbidden", http.StatusForbidden, nil}}
	UnauthorizedResponse        = responseType{false, nil, &errorWrap{"Unauthorized", http.StatusUnauthorized, nil}}
	InternalServerErrorResponse = responseType{false, nil, &errorWrap{"Internal Server Error", http.StatusInternalServerError, nil}}
	BadRequestResponse          = responseType{false, nil, &errorWrap{"Bad Request", http.StatusBadRequest, nil}}
)
