package app

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func (g *AppGin) BindWithParam(v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := g.C.ShouldBind(v)
	if err != nil {
		valierrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for _, value := range valierrs {
			errs = append(errs, &ValidError{
				Key:     value.Namespace(),
				Message: value.Error(),
			})
		}
		return false, errs
	}
	return true, nil
}
