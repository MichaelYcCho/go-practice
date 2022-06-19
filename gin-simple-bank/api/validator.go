package api

import (
	"github.com/MichaelYcCho/go-practice/gin-simple-bank/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currnency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currnency)

	}
	return false
}
