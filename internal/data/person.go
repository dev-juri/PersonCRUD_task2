package data

import (
	"github.com/dev-juri/PersonCRUD_task2/internal/validator"
	"strings"
)

type Person struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Age    int32  `json:"age"`
	Gender string `json:"gender"`
}

func ValidatePerson(v *validator.Validator, person *Person) {

	if person.Age == 0 {
		v.AddError("Age should be greater than 0")
		return
	}

	if person.Age < 0 || strings.TrimSpace(person.Gender) == "" || strings.TrimSpace(person.Name) == "" {
		v.AddError("All fields should be filled.")
		return
	}

	if strings.ToLower(person.Gender) != "male" || strings.ToLower(person.Gender) != "female" {
		v.AddError("Invalid gender specified, should be male of female")
		return
	}
}
