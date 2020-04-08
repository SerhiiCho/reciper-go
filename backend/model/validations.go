package model

import valid "github.com/go-ozzo/ozzo-validation"

// if given condition is true it will validate the field
func requiredIf(cond bool) valid.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return valid.Validate(value, valid.Required)
		}

		return nil
	}
}
