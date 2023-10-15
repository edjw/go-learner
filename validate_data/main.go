// Create a new person struct with name, age, and handedness fields
// Validates the struct using the validator package to make sure the name field is not empty, the age field is between 0 and 130, and the handedness field is either "left", "right", or "ambidextrous"
// Print out the struct and error (if there is one)

package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Person struct {
	Name       string `validate:"required"`
	Age        int    `validate:"gte=0,lte=130"`
	Handedness string `validate:"oneof=left right ambidextrous"`
}

var validate = validator.New()

func logValidationError(person Person, err error) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errors {
			fmt.Printf("%s: Field validation for '%s' failed on the '%s' tag\n", person.Name, e.Field(), e.Tag())
		}
	}
}

func createNewPerson(name string, age int, handed string) (*Person, error) { // returns a pointer to the memory address of a Person struct (or nil if there's an error), and an error

	person := Person{Name: name, Age: age, Handedness: handed}

	err := validate.Struct(person)

	if err != nil {
		logValidationError(person, err)
		return nil, err // return a nil pointer and the error if the data is invalid
	}

	return &person, nil // return the pointer to the memory address holding the person struct and nil if the data is valid

}

// This function will print out the details of a Person only if that Person exists (pPointer != nil) and there's no error (err == nil)
func logPerson(personPointer *Person, err error) {
	if personPointer != nil && err == nil {

		// pPointer is a pointer – a signpost to the actual data. We "dereference" it with *pPointer to follow the signpost to get the actual data. We can do whatever we want with that data including then assign that data to a new variable.
		actualPerson := *personPointer

		// actualPerson is a Person struct not a pointer anymore. We can do whatever we want with that struct including print it out.

		// Do something with the person here if they're successfully created and validated
		fmt.Printf("Created new person: %+v\n", actualPerson)
	}
}

func main() {

	p1, err1 := createNewPerson("Ed", -4, "left")
	p2, err2 := createNewPerson("Erin", 20, "right")
	p3, err3 := createNewPerson("Evan", 18, "xyz")

	logPerson(p1, err1)
	logPerson(p2, err2)
	logPerson(p3, err3)

}
