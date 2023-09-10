package main

import (
	"github.com/dev-juri/PersonCRUD_task2/internal/data"
	"github.com/dev-juri/PersonCRUD_task2/internal/validator"
	"net/http"
)

func (app *application) createPerson(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string `json:"name"`
		Age    int32  `json:"age"`
		Gender string `json:"gender"`
	}

	err := app.readJSON(w, r, input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	person := &data.Person{
		ID:     "1",
		Name:   input.Name,
		Age:    input.Age,
		Gender: input.Gender,
	}

	v := validator.New()

	if data.ValidatePerson(v, person); !v.Valid() {
		app.failedValidationResponse(w, r, v.Error)
		return
	}
	res := result{Status: http.StatusOK, Message: "Person created successfully", Data: nil}

	err = app.writeJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		return
	}
}

func (app *application) fetchPerson(w http.ResponseWriter, r *http.Request) {
	name, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	person := data.Person{
		ID:     "12345678",
		Name:   name,
		Age:    16,
		Gender: "Male",
	}

	res := result{
		Status:  http.StatusOK,
		Message: "Successful",
		Data:    envelope{"person": person},
	}

	err = app.writeJSON(w, http.StatusOK, res, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePerson(w http.ResponseWriter, r *http.Request) {
	_, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	res := result{
		Status:  http.StatusOK,
		Message: "Update Successful",
		Data:    nil,
	}

	err = app.writeJSON(w, http.StatusOK, res, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deletePerson(w http.ResponseWriter, r *http.Request) {
	_, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	res := result{
		Status:  http.StatusOK,
		Message: "Delete Successful",
		Data:    nil,
	}

	err = app.writeJSON(w, http.StatusOK, res, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
