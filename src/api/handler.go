package api

// MultiplyHandler Allows to handle all the REST requests
import (
	"encoding/json"
	"errors"
	"github.com/kenesparta/tkRestService/common"
	"net/http"
)

type MultiplyHandler struct {
}

func (mh *MultiplyHandler) Post(w http.ResponseWriter, r *http.Request) {
	common.CommonHeaders(w)
	common.ValidateHeaders(w, r)
	var (
		multiply Factor
		dec      = json.NewDecoder(r.Body)
	)

	dec.DisallowUnknownFields()
	if err := dec.Decode(&multiply); err != nil {
		common.HttpErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if !multiply.AreValidNumbers() {
		common.HttpErrorResponse(w, http.StatusBadRequest, errors.New("fields required"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&ProductResponse{Result: multiply.Product()}); err != nil {
		common.HttpErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
}
