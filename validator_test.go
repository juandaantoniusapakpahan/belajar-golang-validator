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

func TestMultipleTagError(t *testing.T) {
	var validate *validator.Validate = validator.New()
	data := "salahsalah"

	err := validate.Var(data, "required,numeric")

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestMultipleTagPass(t *testing.T) {

	var validate *validator.Validate = validator.New()
	data := "12423"
	err := validate.Var(data, "required,numeric")

	assert.Equal(t, nil, err)
}

func TestTagParameterError(t *testing.T) {
	var validate *validator.Validate = validator.New()
	data := "1213"

	err := validate.Var(data, "required,numeric,min=5,max=20")

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestTagParameterPass(t *testing.T) {
	var validate *validator.Validate = validator.New()
	data := "134323432289985"
	err := validate.Var(data, "required,numeric,min=5,max=20")
	assert.Equal(t, nil, err)

}

func TestStructTagError(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=6"`
	}

	var validate *validator.Validate = validator.New()

	data := LoginRequest{
		Username: "pastikansalah",
		Password: "sdi23",
	}

	err := validate.Struct(data)

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)

}

func TestStructTagPass(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=6"`
	}

	var validate *validator.Validate = validator.New()

	data := LoginRequest{
		Username: "pastikan93@gmail.com",
		Password: "wer9342",
	}

	err := validate.Struct(data)

	assert.Equal(t, nil, err)

}

func TestValidationError(t *testing.T) {
	var validate *validator.Validate = validator.New()
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=6"`
	}

	data := LoginRequest{
		Username: "error",
		Password: "er13",
	}

	err := validate.Struct(data)

	if err != nil {
		validationError := err.(validator.ValidationErrors)

		for _, fieldError := range validationError {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}
