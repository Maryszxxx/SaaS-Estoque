package validator

import (
	"encoding/json"
	"errors"
	"saas-estoque/entity"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validatorErr error) *entity.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validatorErr, &jsonErr) {
		return entity.NewBadRequestError("Invalid JSON type for field")

	} else if errors.As(validatorErr, &jsonValidationErr) {
		errorsCauses := []entity.Causes{}

		for _, e := range validatorErr.(validator.ValidationErrors) {
			cause := entity.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return entity.NewBadRequestValidationError("Invalid JSON for field ", errorsCauses)

	} else {
		return entity.NewBadRequestError("Invalid JSON body")
	}

}
