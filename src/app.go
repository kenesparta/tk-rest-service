package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kenesparta/tkRestService/api"
	"log"
	"net/http"
)

const addr = ":8084"

type App struct {
	router  *mux.Router
	config  string
	headers handlers.CORSOption
	methods handlers.CORSOption
	origins handlers.CORSOption
}

func (a *App) Initialize() {
	// a.config = configuration.ReadConfiguration("config.toml")
	a.router = mux.NewRouter()
	api.InitRoutes(a.router)
}

func (a *App) Run() {
	a.setCORS()
	if err := http.ListenAndServe(
		addr,
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
