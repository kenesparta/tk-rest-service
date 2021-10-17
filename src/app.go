package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kenesparta/tkRestService/api"
	"log"
	"net/http"
)

const port = ":8084"

type App struct {
	router  *mux.Router
	port    string
	config  string
	headers handlers.CORSOption
	methods handlers.CORSOption
	origins handlers.CORSOption
}

func (a *App) Initialize() {
	a.port = port
	a.router = mux.NewRouter()
	api.InitRoutes(a.router)
}

func (a *App) Run() {
	a.setCORS()
	if err := http.ListenAndServe(
		a.port,
		handlers.CORS(
			a.origins,
			a.headers,
			a.methods,
		)(
			handlers.CompressHandler(a.router),
		),
	); nil != err {
		log.Print(err)
	}
}
