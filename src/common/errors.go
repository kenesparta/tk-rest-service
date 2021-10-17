package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	GenericError = iota
	JSONSyntaxError
	UnexpectedEOFError
	UnmarshalTypeError
	UnknownFieldError
	BodyEmptyError
	BodyTooLargeError
)

type ErrorHandle struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

func HttpErrorResponse(w http.ResponseWriter, err error) {
	var (
		unmarshalTypeError *json.UnmarshalTypeError
		syntaxError        *json.SyntaxError
		eh                 ErrorHandle
	)

	switch {
	case err.Error() == "http: request body too large":
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		eh.Code = BodyTooLargeError
		eh.Message = "Request body must not be larger than 1MB"

	case errors.As(err, &syntaxError):
		w.WriteHeader(http.StatusBadRequest)
		eh.Code = JSONSyntaxError
		eh.Message = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

	case errors.Is(err, io.ErrUnexpectedEOF):
		w.WriteHeader(http.StatusBadRequest)
		eh.Code = UnexpectedEOFError
		eh.Message = fmt.Sprintf("Request body contains badly-formed JSON")

	case errors.As(err, &unmarshalTypeError):
		w.WriteHeader(http.StatusBadRequest)
		eh.Code = UnmarshalTypeError
		eh.Message = fmt.Sprintf(
			"Request body contains an invalid value for the %q field (at position %d)",
			unmarshalTypeError.Field,
			unmarshalTypeError.Offset,
		)

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		w.WriteHeader(http.StatusBadRequest)
		eh.Code = UnknownFieldError
		eh.Message = fmt.Sprintf(
			"Request body contains unknown field %s",
			strings.TrimPrefix(err.Error(), "json: unknown field "),
		)

	case errors.Is(err, io.EOF):
		w.WriteHeader(http.StatusBadRequest)
		eh.Code = BodyEmptyError
		eh.Message = "Request body must not be empty"

	default:
		w.WriteHeader(http.StatusBadRequest)
		eh.Code = GenericError
		eh.Message = err.Error()
	}
	log.Println(eh.Message)
	_ = json.NewEncoder(w).Encode(&eh)
	return
}
