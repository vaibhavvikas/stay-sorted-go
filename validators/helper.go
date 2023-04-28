package validators

import (
	"net/url"

	"github.com/vaibhavvikas/stay-sorted/models"
)

func InvalidRequestValidatorError(err url.Values) error {
	var validationError models.AppError
	for key, fieldErrors := range err {
		validationError.Message = fieldErrors[0]
		validationError.Param = key
	}
	return &validationError
}

func ParamMissingOrInvalidError(param string) error {
	return &models.AppError{
		Param:   param,
		Message: "The param is missign or invalid.",
	}
}

func InvalidRequestError(param, message string) error {
	return &models.AppError{
		Param:   param,
		Message: message,
	}
}
