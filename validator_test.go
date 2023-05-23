package golangvalidator

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()

	assert.NotEqual(t, nil, validate)
}

func TestValidationTagError(t *testing.T) {
	var validate *validator.Validate = validator.New()
	data := ""
	err := validate.Var(data, "required")

	assert.NotEqual(t, nil, err)
	fmt.Println(err.Error())
}

func TestValidationTagCorrect(t *testing.T) {
	var validate *validator.Validate = validator.New()
	data := "tidak kosong bung"

	err := validate.Var(data, "required")

	assert.Equal(t, nil, err)
}

func TestValidationVarWithValueError(t *testing.T) {
	var validate *validator.Validate = validator.New()
	password := "secret"
	confirmPassword := "notsame"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestValidationVarWithValuePass(t *testing.T) {
	var validate *validator.Validate = validator.New()
	password := "secret"
	confirmPassword := "secret"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")

	assert.Equal(t, nil, err)

}