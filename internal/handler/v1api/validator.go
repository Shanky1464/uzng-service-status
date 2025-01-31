package v1api

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("name", ValidateName)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("lastname", ValidateLastName)
	}
}

var reName = regexp.MustCompile(`^[a-zA-Z-]+$`)
var reLastName = regexp.MustCompile(`^[a-zA-Z-]*$`)

var ValidateName validator.Func = func(fl validator.FieldLevel) bool {
	// Only alpha characters and hyphen (-) allowed
	if name, ok := fl.Field().Interface().(string); ok {
		return reName.MatchString(name)
	}
	return false
}

var ValidateLastName validator.Func = func(fl validator.FieldLevel) bool {
	// Nothing or only alpha characters and hyphen (-) allowed
	if lastName, ok := fl.Field().Interface().(string); ok {
		return reLastName.MatchString(lastName)
	}
	return false
}
