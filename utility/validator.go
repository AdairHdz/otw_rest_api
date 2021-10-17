package utility

import (
	"regexp"
	"unicode"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func NewValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("securepass", securePassword)
		validate.RegisterValidation("alpha", alpha)
	}
	return validate
}

func securePassword(fl validator.FieldLevel) bool {
	pass := fl.Field().String()
	containsNumber := false
	containsUpperCase := false
	containsLowerCase := false
	containsSymbol := false

	for _, char := range pass {		
		r := rune(char)
		if unicode.IsNumber(r) {
			containsNumber = true
		}

		if unicode.IsUpper(r) {
			containsUpperCase = true
		}

		if unicode.IsLower(r) {
			containsLowerCase = true
		}

		if unicode.IsSymbol(r) {
			containsSymbol = true
		}		
	}		
	
	return containsNumber && containsLowerCase &&
		containsUpperCase && containsSymbol	
}

func alpha(fl validator.FieldLevel) bool {
	fieldValue := fl.Field().String()
	matches, err := regexp.MatchString(`^[\w'\-,.]*[^_!¡?÷?¿\/\\+=@#$%ˆ&*(){}|~<>;:[\]]*$`, fieldValue)
	if err != nil {
		return false
	}
	return matches
}