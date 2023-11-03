package responses

import "net/http"

type (
	ResponseType struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
		Error   *ErrorWrap  `json:"error"`
	}

	ErrorWrap struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
	}

	defaultErr struct{}
)



var (
	ForbiddenResponse           = ResponseType{false, nil, &ErrorWrap{"Forbidden", http.StatusForbidden, nil}}
	UnauthorizedResponse        = ResponseType{false, nil, &ErrorWrap{"Unauthorized", http.StatusUnauthorized, nil}}
	InternalServerErrorResponse = ResponseType{false, nil, &ErrorWrap{"Internal Server Error", http.StatusInternalServerError, nil}}
	BadRequestResponse          = ResponseType{false, nil, &ErrorWrap{"Bad Request", http.StatusBadRequest, nil}}
)

func ErrorResponse (code int, message string) ResponseType {
	return ResponseType{false, nil, &ErrorWrap{
		Code: code,
		Data: nil,
		Message: message,
	}}
}
