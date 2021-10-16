package api

// MultiplyHandler Allows to handle all the REST requests
import (
	"encoding/json"
	"errors"
	"net/http"
)

type MultiplyHandler struct {
}

func (mh *MultiplyHandler) Post(w http.ResponseWriter, r *http.Request) {
	CommonHeaders(w)
	ValidateHeaders(w, r)
	var (
		multiply Multiply
		dec      = json.NewDecoder(r.Body)
	)

	dec.DisallowUnknownFields()
	if err := dec.Decode(&multiply); err != nil {
		HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if !multiply.AreValidNumbers() {
		HttpErrorResponse(w, http.StatusBadRequest, errors.New("fields required"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&MultiplyResponse{Result: multiply.Result()}); err != nil {
		HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}
