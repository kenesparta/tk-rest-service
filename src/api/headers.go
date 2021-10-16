package api

import (
	"github.com/golang/gddo/httputil/header"
	"net/http"
)

func CommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("X-Frame-Options", "SAMEORIGIN")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func ValidateHeaders(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") == "" {
		http.Error(w, "Content-Type header is required", http.StatusUnsupportedMediaType)
		return
	}
	value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
	if value != "application/json" {
		http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
		return
	}
}

func SetVersion(w http.ResponseWriter, version string) {
	w.Header().Set("Accepts-version", version)
}
