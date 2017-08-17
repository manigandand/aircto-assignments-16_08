package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

type Issue struct {
	Title       string `validate:"required"`
	Description string `validate:"required"`
	AssignedTo  int    `validate:"required"`
	Status      string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {

	validate = validator.New()

	validateStruct()
	validateVariable()
}

func validateStruct() {

	user := &Issue{
		Title:       "Eavesdown Docks",
		Description: "",
		AssignedTo:  2,
		Status:      "open",
	}

	// returns nil or ValidationErrors ( map[string]*FieldError )
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
		var erors []string
		for _, err := range err.(validator.ValidationErrors) {
			ier := err.Field() + " filed is " + err.Tag()
			// fmt.Println(ier)
			// fmt.Println()
			erors = append(erors, ier)
		}

		// from here you can create your own error messages in whatever language you wish
		fmt.Println(erors)
		return
	}

	// save user to database
}

func validateVariable() {

	myEmail := ""

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Printf("Err(s):\n%+v\n", errs)
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
