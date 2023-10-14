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

func createNewPerson(name string, age int, handed string) (*Person, error) { // returns a pointer to a Person struct and an error

	person := Person{Name: name, Age: age, Handedness: handed}

	err := validate.Struct(person)

	if err != nil {
		logValidationError(person, err)
		return nil, err // return a nil pointer and the error if the data is invalid
	}

	return &person, nil // return a pointer to the person struct if the data is valid

	// a pointer is a memory address that points to the location of a value in memory
	// a value is the actual data stored in memory

}

func logPerson(p *Person, err error) {
	if p != nil && err == nil {
		p := *p // dereference the pointer to get the actual value

		// Do something with the person here if they're successfully created and validated
		fmt.Printf("Created new person: %+v\n", p)
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
