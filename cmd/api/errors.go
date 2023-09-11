package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string) {
	res := result{
		Status:  int32(status),
		Message: message,
		Data:    nil,
	}

	err := app.writeJSON(w, status, res, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, error string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, error)
}

func (app *application) personExists(w http.ResponseWriter, r *http.Request, name string) {
	message := fmt.Sprintf("the person with name %s already exists", name)
	app.errorResponse(w, r, http.StatusUnprocessableEntity, message)
}

func (app *application) personNotFound(w http.ResponseWriter, r *http.Request, name string) {
	message := fmt.Sprintf("the person with name %s doesn't exists", name)
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) updateFailed(w http.ResponseWriter, r *http.Request, name string) {
	message := fmt.Sprintf("update failed, please confirm %s is an existing user", name)
	app.errorResponse(w, r, http.StatusNotFound, message)
}
