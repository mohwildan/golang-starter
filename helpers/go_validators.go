package helpers

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	cvalidator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type CValidator struct {
	Validator *validator.Validate
	Trans     *ut.Translator
}

func RegisterCustomTranslations(v *validator.Validate, trans ut.Translator) {
	// translation
	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "Tidak Boleh Kosong", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	// ... Add other translation registration functions here ...
}

func RegisterCustomValidations(v *validator.Validate) {
	// validation
	_ = v.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		secure := true
		tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
		for _, test := range tests {
			t, _ := regexp.MatchString(test, fl.Field().String())
			if !t {
				secure = false
				break
			}
		}
		return secure
	})
	_ = v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		secure := true

		rgx := regexp.MustCompile(`((0|\+62|062|62)[0-9]{9,14}$)`)
		if !rgx.MatchString(fl.Field().String()) {
			secure = false
		}
		return secure
	})
	// ... Add other validation registration functions here ...
}

func InitValidator() CValidator {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("Data Not Found translator")
		panic("Data Not Found translator")
	}
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}

	RegisterCustomTranslations(v, trans)
	RegisterCustomValidations(v)

	return CValidator{
		Validator: v,
		Trans:     &trans,
	}
}
func ValidateStruct(model interface{}) map[string]interface{} {
	validation := make(map[string]interface{})
	initValidator := InitValidator()
	err := initValidator.Validator.Struct(model)
	if err != nil {
		fmt.Println(err)
		for _, e := range err.(cvalidator.ValidationErrors) {
			fieldName := strings.ReplaceAll(e.Field(), "", "")
			validation[fieldName] = e.Translate(*initValidator.Trans)
		}
	}
	return validation
}
