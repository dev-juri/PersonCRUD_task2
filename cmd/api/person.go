package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dev-juri/PersonCRUD_task2/internal/data"
	"github.com/dev-juri/PersonCRUD_task2/internal/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (app *application) getDatabase() *mongo.Database {
	return app.dbClient.Database("person-db")
}

func (app *application) createPerson(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	person := &data.Person{
		Name:   input.Name,
		Age:    input.Age,
		Gender: input.Gender,
	}

	v := validator.New()
	if data.ValidatePerson(v, person); !v.Valid() {
		app.failedValidationResponse(w, r, v.Error)
		return
	}

	p := app.getDatabase().Collection("person")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err = p.InsertOne(ctx, bson.M{
		"_id":    primitive.NewObjectID(),
		"name":   person.Name,
		"age":    person.Age,
		"gender": person.Gender,
	})

	if err != nil {
		app.serverErrorResponse(w, r, err)
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

	p := app.getDatabase().Collection("person")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.FindOne()

	var person data.Person
	err = p.FindOne(ctx, bson.M{"name": name}, opts).Decode(&person)
	if err != nil {
		app.personNotFound(w, r, name)
		return
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
	name, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	p := app.getDatabase().Collection("person")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var person data.Person
	err = p.FindOne(ctx, bson.M{"name": name}, options.FindOne()).Decode(&person)
	if err != nil {
		app.personNotFound(w, r, name)
		return
	}

	u := bson.D{{"$set", bson.D{{"name", input.Name}, {"age", input.Age}, {"gender", input.Gender}}}}

	_, err = p.UpdateOne(ctx, bson.D{{"_id", person.ID}}, u, options.Update())
	if err != nil {
		app.updateFailed(w, r, name)
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
	name, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	p := app.getDatabase().Collection("person")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err = p.DeleteOne(ctx, bson.M{"name": name}, options.Delete())

	if err != nil {
		app.serverErrorResponse(w, r, errors.New("internal server error"))
		return
	}

	result := result{
		Status:  http.StatusOK,
		Message: "Delete Successful",
		Data:    nil,
	}

	err = app.writeJSON(w, http.StatusOK, result, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
