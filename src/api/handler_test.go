package api

import (
	"bytes"
	"encoding/json"
	"github.com/kenesparta/tkRestService/common"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMultiplyHandler_Post_ReturnStatusOk(t *testing.T) {
	var (
		jsonStr = []byte(`{"first_factor":4,"second_factor":5}`)
		mh      MultiplyHandler
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMultiplyHandler_Post_ReturnJSONSyntaxError(t *testing.T) {
	var (
		jsonStr = []byte(`{"first_factor":"4,"second_factor":abcd}`)
		mh      MultiplyHandler
		eh      common.ErrorHandle
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&eh); err != nil {
		t.Fatal(err)
	}

	if common.JSONSyntaxError != eh.Code {
		t.Errorf("The error not correspond to %v", eh.Code)
	}
}

func TestMultiplyHandler_Post_ReturnUnexpectedEOFError(t *testing.T) {
	var (
		jsonStr = []byte(`{"first_factor":"4","second_factor":"abcd}`)
		mh      MultiplyHandler
		eh      common.ErrorHandle
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&eh); err != nil {
		t.Fatal(err)
	}

	if common.UnexpectedEOFError != eh.Code {
		t.Errorf("Type error doesn't match")
	}
}

func TestMultiplyHandler_Post_ReturnUnmarshalTypeError(t *testing.T) {
	var (
		jsonStr = []byte(`{"first_factor":4,"second_factor":"abcd"}`)
		mh      MultiplyHandler
		eh      common.ErrorHandle
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&eh); err != nil {
		t.Fatal(err)
	}

	if common.UnmarshalTypeError != eh.Code {
		t.Errorf("Type error doesn't match")
	}
}

func TestMultiplyHandler_Post_ReturnUnknownFieldError(t *testing.T) {
	var (
		jsonStr = []byte(`{"45_first_factor":4.28,"second_factor_factor":15.25}`)
		mh      MultiplyHandler
		eh      common.ErrorHandle
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&eh); err != nil {
		t.Fatal(err)
	}

	if common.UnknownFieldError != eh.Code {
		t.Errorf("Type error doesn't match")
	}
}

func TestMultiplyHandler_Post_ReturnBodyEmptyError(t *testing.T) {
	var (
		jsonStr = []byte(``)
		mh      MultiplyHandler
		eh      common.ErrorHandle
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&eh); err != nil {
		t.Fatal(err)
	}

	if common.BodyEmptyError != eh.Code {
		t.Errorf("Type error doesn't match")
	}
}

func TestMultiplyHandler_Post_ReturnFieldsRequired(t *testing.T) {
	var (
		jsonStr = []byte(`{"first_factor":4.28}`)
		mh      MultiplyHandler
		eh      common.ErrorHandle
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&eh); err != nil {
		t.Fatal(err)
	}

	if "fields required" != eh.Message && eh.Code == common.GenericError {
		t.Errorf("Type error doesn't match")
	}
}

func TestMultiplyHandler_Post_ReturnProductResponse(t *testing.T) {
	var (
		jsonStr = []byte(`{"first_factor":5.147, "second_factor":48.111}`)
		mh      MultiplyHandler
		pr      ProductResponse
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&pr); err != nil {
		t.Fatal(err)
	}

	if 247.627 != pr.Product {
		t.Errorf("Calculation error")
	}
}

func TestMultiplyHandler_Post_ReturnInfiniteProduct(t *testing.T) {
	var (
		jsonStr = []byte(`{"first_factor":-2222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222.12, "second_factor":-1522222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222.5}`)
		mh      MultiplyHandler
		eh      common.ErrorHandle
	)

	req, err := http.NewRequest("POST", "/multiply", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mh.Post)
	handler.ServeHTTP(rr, req)

	if err := json.NewDecoder(rr.Body).Decode(&eh); err != nil {
		t.Fatal(err)
	}

	if "infinite product" != eh.Message && eh.Code == common.GenericError {
		t.Errorf("Type error doesn't match")
	}
}