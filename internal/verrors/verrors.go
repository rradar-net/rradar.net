package verrors

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var trans ut.Translator

func Register() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Use the names which have been specified for JSON representations of structs, rather than normal Go field names.
		// eg. "username" instead of "Username"
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

		en := en.New()
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(v, trans)
	}
}

func Format(err error) []string {
	ve := err.(validator.ValidationErrors)
	errs := make([]string, len(ve))
	for i, e := range ve {
		errs[i] = e.Translate(trans)
	}
	return errs
}
