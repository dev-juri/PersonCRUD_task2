package main

import (
	"context"
	"errors"
	"fmt"
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

	value, err := p.InsertOne(ctx, bson.M{
		"_id":    primitive.NewObjectID(),
		"name":   person.Name,
		"age":    person.Age,
		"gender": person.Gender,
	})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	person.ID = value.InsertedID

	res := result{Status: http.StatusOK, Message: "Person created successfully", Data: envelope{"person": person}}

	err = app.writeJSON(w, http.StatusCreated, res, nil)
	if err != nil {
		return
	}
}

func (app *application) fetchPerson(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	p := app.getDatabase().Collection("person")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.FindOne()

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		app.serverErrorResponse(w, r, mongo.ErrInvalidIndexValue)
		return
	}

	var person data.Person
	err = p.FindOne(ctx, bson.M{"_id": _id}, opts).Decode(&person)
	if err != nil {
		app.personNotFound(w, r, id)
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
	id, err := app.readIDParam(r)

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

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		app.serverErrorResponse(w, r, mongo.ErrInvalidIndexValue)
		return
	}

	p := app.getDatabase().Collection("person")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var person data.Person

	u := bson.D{{"$set", bson.D{{"name", input.Name}, {"age", input.Age}, {"gender", input.Gender}}}}

	_, err = p.UpdateOne(ctx, bson.D{{"_id", _id}}, u, options.Update())
	if err != nil {
		app.updateFailed(w, r, id)
		return
	}

	person.ID = id
	person.Name = input.Name
	person.Age = input.Age
	person.Gender = input.Gender

	res := result{
		Status:  http.StatusOK,
		Message: "Update Successful",
		Data:    envelope{"person": person},
	}

	err = app.writeJSON(w, http.StatusOK, res, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deletePerson(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		app.serverErrorResponse(w, r, mongo.ErrInvalidIndexValue)
		return
	}

	p := app.getDatabase().Collection("person")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err = p.DeleteOne(ctx, bson.M{"_id": _id}, options.Delete())

	if err != nil {
		app.serverErrorResponse(w, r, errors.New("internal server error"))
		return
	}

	result := result{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("Record for person with id %s deleted successfully", id),
		Data:    nil,
	}

	err = app.writeJSON(w, http.StatusOK, result, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
