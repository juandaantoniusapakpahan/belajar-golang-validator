package golangvalidator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func TestStrucCrossFieldError(t *testing.T) {
	var validate *validator.Validate = validator.New()
	type LoginRequest struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=6"`
		ConfirmPassword string `validate:"required,min=6,eqfield=Password"`
	}

	data := LoginRequest{
		Username:        "gimai@gmail.com",
		Password:        "3e423dsa",
		ConfirmPassword: "awefa903",
	}

	err := validate.Struct(data)
	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestStrucCrossFieldPass(t *testing.T) {
	var validate *validator.Validate = validator.New()
	type LoginRequest struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=6"`
		ConfirmPassword string `validate:"required,min=6,eqfield=Password"`
	}

	data := LoginRequest{
		Username:        "gimai@gmail.com",
		Password:        "3e423dsa",
		ConfirmPassword: "3e423dsa",
	}

	err := validate.Struct(data)
	assert.Equal(t, nil, err)
}

func TestStructNestedError(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}
	type Person struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Address: Address{
			Street: "",
			City:   "",
		},
	}

	err := validate.Struct(data)

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestStructNestedPass(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}
	type Person struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Address: Address{
			Street: "Jalan CutNyadien",
			City:   "Jakarta Selatan",
		},
	}

	err := validate.Struct(data)

	assert.Equal(t, nil, err)
}

// use "dive" tag
func TestCollectionStructNestedError(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}
	type Person struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"` // dive tag for collection
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "",
				City:   "",
			},
			{
				Street: "",
				City:   "",
			},
		},
	}

	err := validate.Struct(data)
	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestCollectionStructNestedPass(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}
	type Person struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"` // dive tag for collection
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "Jln Djuanda",
				City:   "Jakarta selatan",
			},
			{
				Street: "Jln Gatot",
				City:   "Jakarta Pusat",
			},
		},
	}

	err := validate.Struct(data)
	assert.Equal(t, nil, err)
}

func TestBasicCollectionError(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}
	type Person struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"` // dive tag for collection
		Hobbies   []string  `validate:"required,dive,required,min=4"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "Jln Djuanda",
				City:   "Jakarta selatan",
			},
			{
				Street: "Jln Gatot",
				City:   "Jakarta Pusat",
			},
		},
		Hobbies: []string{
			"makan",
			"mabuk",
			"",
			"tdr",
		},
	}

	err := validate.Struct(data)
	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}
func TestBasicCollectionPass(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}
	type Person struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"` // dive tag for collection
		Hobbies   []string  `validate:"required,dive,required,min=4"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "Jln Djuanda",
				City:   "Jakarta selatan",
			},
			{
				Street: "Jln Gatot",
				City:   "Jakarta Pusat",
			},
		},
		Hobbies: []string{
			"makan",
			"mabuk",
			"traveling",
			"tidur",
		},
	}

	err := validate.Struct(data)
	assert.Equal(t, nil, err)
}

func TestMapValidationError(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type Person struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"` // dive tag for collection
		Hobbies   []string          `validate:"required,dive,required,min=4"`
		School    map[string]School `validate:"required,dive,keys,required,min=3,endkeys,required"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "Jln Djuanda",
				City:   "Jakarta selatan",
			},
			{
				Street: "Jln Gatot",
				City:   "Jakarta Pusat",
			},
		},
		Hobbies: []string{
			"makan",
			"mabuk",
			"dfasdf",
			"tsfsdfdr",
		},
		School: map[string]School{
			"DSFW": {Name: "GGWP"},
			"DWDF": {Name: ""},
			"":     {Name: "werwae"},
			"ds":   {Name: "we98we"},
		},
	}

	err := validate.Struct(data)
	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestMapValidationPass(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type Person struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"` // dive tag for collection
		Hobbies   []string          `validate:"required,dive,required,min=4"`
		School    map[string]School `validate:"required,dive,keys,required,min=3,endkeys,required"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "Jln Djuanda",
				City:   "Jakarta selatan",
			},
			{
				Street: "Jln Gatot",
				City:   "Jakarta Pusat",
			},
		},
		Hobbies: []string{
			"makan",
			"mabuk",
			"dfasdf",
			"tsfsdfdr",
		},
		School: map[string]School{
			"DSFW": {Name: "GGWP"},
			"DWDF": {Name: "asdfsa"},
			"DSFS": {Name: "werwae"},
			"DSFV": {Name: "we98we"},
		},
	}

	err := validate.Struct(data)
	assert.Equal(t, nil, err)
}

func TestMapValidationBasicError(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type Person struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"` // dive tag for collection
		Hobbies   []string          `validate:"required,dive,required,min=4"`
		School    map[string]School `validate:"required,dive,keys,required,min=3,endkeys,required"`
		Item      map[string]int    `validate:"required,dive,keys,required,endkeys,required,gt=100"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "Jln Djuanda",
				City:   "Jakarta selatan",
			},
			{
				Street: "Jln Gatot",
				City:   "Jakarta Pusat",
			},
		},
		Hobbies: []string{
			"makan",
			"mabuk",
			"dfasdf",
			"tsfsdfdr",
		},
		School: map[string]School{
			"DSFW": {Name: "GGWP"},
			"DWDF": {Name: "asdfsa"},
			"DSFS": {Name: "werwae"},
			"DSFV": {Name: "we98we"},
		},
		Item: map[string]int{
			"A": 234,
			"":  23423,
			"C": 32,
		},
	}

	err := validate.Struct(data)
	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestMapValidationBasicPass(t *testing.T) {
	type Address struct {
		Street string `validate:"required"`
		City   string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type Person struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"` // dive tag for collection
		Hobbies   []string          `validate:"required,dive,required,min=4"`
		School    map[string]School `validate:"required,dive,keys,required,min=3,endkeys,required"`
		Item      map[string]int    `validate:"required,dive,keys,required,endkeys,required,gt=100"`
	}

	var validate *validator.Validate = validator.New()

	data := Person{
		Id:   "dasij23423",
		Name: "GGwps",
		Addresses: []Address{
			{
				Street: "Jln Djuanda",
				City:   "Jakarta selatan",
			},
			{
				Street: "Jln Gatot",
				City:   "Jakarta Pusat",
			},
		},
		Hobbies: []string{
			"makan",
			"mabuk",
			"dfasdf",
			"tsfsdfdr",
		},
		School: map[string]School{
			"DSFW": {Name: "GGWP"},
			"DWDF": {Name: "asdfsa"},
			"DSFS": {Name: "werwae"},
			"DSFV": {Name: "we98we"},
		},
		Item: map[string]int{
			"A": 234,
			"B": 23423,
			"C": 3232,
		},
	}

	err := validate.Struct(data)
	assert.Equal(t, nil, err)
}

func TestAliasError(t *testing.T) {
	var validate *validator.Validate = validator.New()
	validate.RegisterAlias("varchar", "min=5")

	type Profile struct {
		Id    string `validate:"varchar"`
		Name  string `validate:"varchar"`
		Email string `validate:"varchar"`
	}

	data := Profile{
		Id:    "dsaf",
		Name:  "iasd",
		Email: "fifaowefj",
	}

	err := validate.Struct(data)

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestAliasPass(t *testing.T) {
	var validate *validator.Validate = validator.New()
	validate.RegisterAlias("varchar", "min=5")

	type Profile struct {
		Id    string `validate:"varchar"`
		Name  string `validate:"varchar"`
		Email string `validate:"varchar"`
	}

	data := Profile{
		Id:    "dsasdff",
		Name:  "iasdfsd",
		Email: "fifaowefj",
	}

	err := validate.Struct(data)

	assert.Equal(t, nil, err)
}

func MustValidateUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}

func TestFucnValidateError(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidateUsername)

	type User struct {
		Username string `validate:"username"`
		Password string `validate:"required"`
	}

	data := User{
		Username: "asfasSDF",
		Password: "faweijwa",
	}

	err := validate.Struct(data)

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}

func TestFucnValidatePass(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidateUsername)

	type User struct {
		Username string `validate:"username"`
		Password string `validate:"required"`
	}

	data := User{
		Username: "DJFSD",
		Password: "faweijwa",
	}

	err := validate.Struct(data)

	assert.Equal(t, nil, err)
}

var regestNumber = regexp.MustCompile("^[0-9]+$")

func MustCustomValidate(field validator.FieldLevel) bool {
	num, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}

	value := field.Field().String()

	if !regestNumber.MatchString(value) {
		return false
	}

	return len(value) == num
}

func TestCustomValidationError(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustCustomValidate)

	type DataTest struct {
		Pin      string `validate:"pin=5"`
		Username string `validate:"required"`
	}

	data := DataTest{
		Pin:      "32432",
		Username: "asdfasd",
	}

	err := validate.Struct(data)

	fmt.Println(err.Error())
	assert.NotEqual(t, nil, err)
}
func TestCustomValidationPass(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustCustomValidate)

	type DataTest struct {
		Pin      string `validate:"pin=5"`
		Username string `validate:"required"`
	}

	data := DataTest{
		Pin:      "32423",
		Username: "asdfasd",
	}

	err := validate.Struct(data)

	assert.Equal(t, nil, err)
}

func TestOrValidation(t *testing.T) {
	validate := validator.New()

	type DataTest struct {
		Username string `validate:"email|numeric"`
		Password string `validate:"required"`
	}

	data1 := DataTest{
		Username: "asdjf@gmail.com",
		Password: "sasfas",
	}
	data2 := DataTest{
		Username: "23432423",
		Password: "asdfiaosdf",
	}

	err := validate.Struct(data1)
	assert.Equal(t, nil, err)

	err = validate.Struct(data2)
	assert.Equal(t, nil, err)
}
