package common

import (
	"errors"
	"github.com/golang/gddo/httputil/header"
	"net/http"
)

func Headers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func ValidateHeaders(r *http.Request) error {
	if r.Header.Get("Content-Type") == "" {
		return errors.New("Content-Type header is required")
	}
	value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
	if value != "application/json" {
		return errors.New("Content-Type header is not application/json")
	}
	return nil
}
