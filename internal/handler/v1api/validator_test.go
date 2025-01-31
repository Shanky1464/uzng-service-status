//go:build unit

package v1api_test

import (
	"strconv"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"uniphore.com/platform-hello-world-go/internal/handler/v1api"
)

func TestValidateName(t *testing.T) {
	v := validator.New()
	v.RegisterValidation("name", v1api.ValidateName)

	validName := "FooBar"
	err := v.Var(validName, "name")
	assert.NoError(t, err, "Should validate name correctly")
}

func TestValidateIncorrectName(t *testing.T) {
	cases := []interface{}{
		"FooBar1234",
		1234,
		true,
		"",
	}

	v := validator.New()
	v.RegisterValidation("name", v1api.ValidateName)

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := v.Var(tc, "name")
			assert.Error(t, err, "Should not validate name correctly")
		})
	}
}

func TestValidateLastName(t *testing.T) {
	cases := []string{
		"FooBar",
		"",
	}

	v := validator.New()
	v.RegisterValidation("lastname", v1api.ValidateLastName)

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := v.Var(tc, "lastname")
			assert.NoError(t, err, "Should validate last name correctly")
		})
	}
}

func TestValidateIncorrectLastName(t *testing.T) {
	cases := []interface{}{
		"FooBar1234",
		1234,
		true,
	}

	v := validator.New()
	v.RegisterValidation("lastname", v1api.ValidateLastName)

	for i, tc := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			err := v.Var(tc, "lastname")
			assert.Error(t, err, "Should not validate last name correctly")
		})
	}
}
