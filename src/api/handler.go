package api

import (
	"encoding/json"
	"errors"
	"github.com/kenesparta/multiplyLogic"
	"github.com/kenesparta/tkRestService/common"
	"net/http"
)

// MultiplyHandler Allows to handle all the REST requests
type MultiplyHandler struct{}

func (mh *MultiplyHandler) Post(w http.ResponseWriter, r *http.Request) {
	common.Headers(w)
	if err := common.ValidateHeaders(r); err != nil {
		common.HttpErrorResponse(w, err)
		return
	}
	var (
		multiply multiplyLogic.Factor
		dec      = json.NewDecoder(r.Body)
	)

	dec.DisallowUnknownFields()
	if err := dec.Decode(&multiply); err != nil {
		common.HttpErrorResponse(w, err)
		return
	}

	if !multiply.AreValidNumbers() {
		common.HttpErrorResponse(w, errors.New("fields required"))
		return
	}

	if multiply.IsProductInfinite() {
		common.HttpErrorResponse(w, errors.New("infinite product"))
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&ProductResponse{Product: multiply.Product()})
}
