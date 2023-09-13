package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodPost, "/api", app.createPerson)
	router.HandlerFunc(http.MethodGet, "/api/:id", app.fetchPerson)
	router.HandlerFunc(http.MethodPut, "/api/:id", app.updatePerson)
	router.HandlerFunc(http.MethodDelete, "/api/:id", app.deletePerson)

	return router
}
