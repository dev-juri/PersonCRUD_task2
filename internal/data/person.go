package data

import (
	"strings"

	"github.com/dev-juri/PersonCRUD_task2/internal/validator"
)

type Person struct {
	ID     interface{} `json:"id,omitempty" bson:"_id"`
	Name   string      `json:"name" bson:"name"`
	Age    int         `json:"age" bson:"age"`
	Gender string      `json:"gender" bson:"gender"`
}

func ValidatePerson(v *validator.Validator, person *Person) {

	if person.Age <= 0 {
		v.AddError("Age should be greater than 0")
		return
	}

	if person.Age < 0 || strings.TrimSpace(person.Gender) == "" || strings.TrimSpace(person.Name) == "" {
		v.AddError("All fields should be filled.")
		return
	}

	/***if strings.ToLower(person.Gender) != "male" || strings.ToLower(person.Gender) != "female" {
		v.AddError("gender should be male or female")
		return
	}***/
}
