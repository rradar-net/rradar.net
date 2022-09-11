package errs

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rradar-net/rradar.net/pkg/proto"
)

var trans ut.Translator

func Init() {
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

func Format(err error) proto.ErrorResponse {
	var ve validator.ValidationErrors

	if !errors.As(err, &ve) {
		msg := "invalid input"
		return proto.ErrorResponse{
			Status:  proto.Status_error,
			Message: &msg,
		}
	}

	errs := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errs[e.Field()] = e.Translate(trans)
	}
	return proto.ErrorResponse{
		Status: proto.Status_fail,
		Data:   errs,
	}
}
