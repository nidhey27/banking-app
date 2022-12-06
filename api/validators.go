package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/nidhey27/banking-app/utils"
)

var validateCurrency validator.Func = func(filedLevel validator.FieldLevel) bool {
	if currency, ok := filedLevel.Field().Interface().(string); ok {
		// Check currency is supported or not
		return utils.IsSupportedCurrency(currency)
	}

	return false
}
