package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validatorEn "github.com/go-playground/validator/v10/translations/en"
	"github.com/rodrigoachilles/auction-go/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enTranslator := en.New()
		enTransl := ut.New(enTranslator, enTranslator)
		transl, _ = enTransl.GetTranslator("en")
		_ = validatorEn.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewNotFoundError("Invalid type error")
	} else if errors.As(validationErr, &jsonValidation) {
		var errorCauses []rest_err.Causes

		for _, e := range validationErr.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}

		return rest_err.NewBadRequestError("Invalid field values", errorCauses...)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
