package response

import (
	"errors"
	"net/http"
)

// error general
var (
	ErrNotFound = errors.New("not found")
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrEmailInvalid     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordInvalid  = errors.New("password must have minimum 6 characters")
	ErrAuthIsNotExists  = errors.New("auth is not exists")
	ErrEmailAlreadyUsed = errors.New("email already used")
	ErrPasswordNotMatch = errors.New("password not match")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral    = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest = NewError("bad request", "40000", http.StatusBadRequest)
)

var (
	ErrorEmailRequired    = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid     = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalid  = NewError(ErrPasswordInvalid.Error(), "40004", http.StatusBadRequest)
	ErrorAuthIsNotExists  = NewError(ErrAuthIsNotExists.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyUsed = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
	ErrorNotFound         = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():         ErrorNotFound,
		ErrEmailRequired.Error():    ErrorEmailRequired,
		ErrEmailInvalid.Error():     ErrorEmailInvalid,
		ErrPasswordRequired.Error(): ErrorPasswordRequired,
		ErrPasswordInvalid.Error():  ErrorPasswordInvalid,
		ErrAuthIsNotExists.Error():  ErrorAuthIsNotExists,
		ErrEmailAlreadyUsed.Error(): ErrorEmailAlreadyUsed,
		ErrPasswordNotMatch.Error(): ErrorPasswordNotMatch,
	}
)
