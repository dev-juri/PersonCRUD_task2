package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodPost, "/api/person", app.createPerson)
	router.HandlerFunc(http.MethodGet, "/api/person/:name", app.fetchPerson)
	router.HandlerFunc(http.MethodPatch, "/api/update/:name", app.updatePerson)
	router.HandlerFunc(http.MethodDelete, "/api/delete/:name", app.deletePerson)

	return router
}
