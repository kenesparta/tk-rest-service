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
	SyntaxError
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

func HttpErrorResponse(w http.ResponseWriter, statusHttpError int, err error) {
	var (
		unmarshalTypeError *json.UnmarshalTypeError
		syntaxError *json.SyntaxError
		em          ErrorHandle
	)

	switch {
	case errors.As(err, &syntaxError):
		w.WriteHeader(http.StatusBadRequest)
		em.Code = SyntaxError
		em.Message = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

	case errors.Is(err, io.ErrUnexpectedEOF):
		w.WriteHeader(http.StatusBadRequest)
		em.Code = UnexpectedEOFError
		em.Message = fmt.Sprintf("Request body contains badly-formed JSON")

	case errors.As(err, &unmarshalTypeError):
		w.WriteHeader(http.StatusBadRequest)
		em.Code = UnmarshalTypeError
		em.Message = fmt.Sprintf(
			"Request body contains an invalid value for the %q field (at position %d)",
			unmarshalTypeError.Field,
			unmarshalTypeError.Offset,
		)

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		w.WriteHeader(http.StatusBadRequest)
		em.Code = UnknownFieldError
		em.Message = fmt.Sprintf(
			"Request body contains unknown field %s",
			strings.TrimPrefix(err.Error(), "json: unknown field "),
		)

	case errors.Is(err, io.EOF):
		w.WriteHeader(http.StatusBadRequest)
		em.Code = BodyEmptyError
		em.Message = "Request body must not be empty"

	case err.Error() == "http: request body too large":
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		em.Code = BodyTooLargeError
		em.Message = "Request body must not be larger than 1MB"

	default:
		w.WriteHeader(http.StatusBadRequest)
		em.Code = GenericError
		em.Message = err.Error()
	}
	log.Println(em.Message)
	_ = json.NewEncoder(w).Encode(&em)
	return
}
