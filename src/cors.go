package main

import "github.com/gorilla/handlers"

func (a *App) setCORS() {
	a.headers = handlers.AllowedHeaders([]string{
		"X-Request-With",
		"Content-Type",
		"Accept",
		"Authorization",
		"Origin",
	})

	a.methods = handlers.AllowedMethods([]string{
		"GET",
		"POST",
		"PUT",
		"PATCH",
		"DELETE",
		"HEAD",
		"OPTIONS",
	})

	a.origins = handlers.AllowedOrigins([]string{
		"*",
	})
}
