package response

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Response struct {
	Status string `json:"status"`
	Error  error  `json:"error,omitempty"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{Status: StatusOk}
}

func Error(msg string) Response {
	return Response{Status: StatusError, Error: errors.New(msg)}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is not a valid URL field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("%s is not a valid field %s", err.Tag(), err.Field()))
		}
	}
	return Response{
		Status: StatusError,
		Error:  errors.New(strings.Join(errMsgs, ";")),
	}
}
