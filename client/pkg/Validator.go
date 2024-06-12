package pkg

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateInputs(data interface{}) (bool, map[string]string) {
	// var validate *validator.Validate

	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(data)
	if err != nil {

		// Validation Syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		// Validation Errors
		errors := make(map[string]string)

		reflected := reflect.ValueOf(data)

		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())

			var name string

			if name = field.Tag.Get("json"); name != "" {
				name = strings.ToLower(err.StructField())
			}

			switch err.Tag() {
			case "required":
				errors[name] = "The " + name + " is required"
				break
			case "email":
				errors[name] = "The " + name + " should be a valid email"
				break
			case "min":
				errors[name] = name + " must be a least " + err.Param() + " characters"
				break
			case "max":
				errors[name] = name + " must be a at most " + err.Param() + " characters"

			default:
				errors[name] = "The " + name + "is invalid"
				break
			}
		}

		return false, errors
	}
	return true, nil
}

func ValidationError(w http.ResponseWriter, HttpCode int, errors map[string]string) {
	response, err := json.Marshal(errors)
	if err != nil {
		log.Print("failed to encode json %w", err)
		http.Error(w, "An Error Occurred", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HttpCode)
	w.Write(response)
}
