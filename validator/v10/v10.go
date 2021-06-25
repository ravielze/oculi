package v10

import (
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ravielze/oculi/constant/errors"
	stderr "github.com/ravielze/oculi/standard/errors"
	"github.com/ravielze/oculi/validator"
)

type (
	impl struct {
		instance *v10.Validate
		trans    *ut.Translator
	}
)

var once sync.Once
var instance validator.Validator = nil

func Instance() (validator.Validator, error) {
	if instance == nil {
		var transOuter *ut.Translator
		var vOuter *v10.Validate
		var errOuter error = nil
		once.Do(func() {
			transEn := en.New()
			transId := id.New()
			uni := ut.New(transEn, transEn, transId)
			trans, _ := uni.GetTranslator("en")
			errOuter = nil
			transOuter = &trans
			if v, ok := binding.Validator.Engine().(*v10.Validate); ok {
				if err := en_translations.RegisterDefaultTranslations(v, trans); err != nil {
					errOuter = err
					return
				}
				vOuter = v
			} else {
				errOuter = errors.ErrValidatorEngineNotFound
			}
		})
		if errOuter != nil {
			once = sync.Once{}
			instance = nil
			return nil, errOuter
		}
		instance = &impl{instance: vOuter, trans: transOuter}
	}
	return instance, nil
}

func (i *impl) RegisterAlias(alias string, tags string) {
	i.instance.RegisterAlias(alias, tags)
}

func (i *impl) RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...interface{}) {
	i.instance.RegisterCustomTypeFunc(v10.CustomTypeFunc(fn), types...)
}

func (i *impl) RegisterValidation(tag string, fn interface{}) {
	if fnConv, ok := fn.(v10.Func); !ok {
		panic("fn is not v10.Func")
	} else {
		i.instance.RegisterValidation(tag, fnConv)
	}
}

func (i *impl) RegisterStructValidation(fn interface{}, types ...interface{}) {
	if fnConv, ok := fn.(v10.StructLevelFunc); !ok {
		panic("fn is not v10.StructLevelFunc")
	} else {
		i.instance.RegisterStructValidation(fnConv, types...)
	}
}

func (i *impl) Validate(object interface{}) error {
	return i.TranslateError(i.instance.Struct(object))
}

func (i *impl) RegisterTranslation(tag string, registerFn interface{}, transFn interface{}) error {
	if registerFnConv, ok := registerFn.(v10.RegisterTranslationsFunc); !ok {
		panic("registerFn is not v10.RegisterTranslationsFunc")
	} else if transFnConv, ok2 := transFn.(v10.TranslationFunc); !ok2 {
		panic("transFn is not v10.TranslationsFunc")
	} else if ok && ok2 {
		return i.instance.RegisterTranslation(tag, *i.trans, registerFnConv, transFnConv)
	} else {
		return errors.ErrUnclasified
	}
}

func (i *impl) AddTranslation(tag string, errorMsg string) error {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errorMsg, false)
	}
	transFn := func(ut ut.Translator, fe v10.FieldError) string {
		t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}
	return i.instance.RegisterTranslation(tag, *i.trans, registerFn, transFn)
}

func (i *impl) TranslateError(err error) error {
	if err == nil {
		return nil
	}
	validatorErrs := err.(v10.ValidationErrors)
	translatedErrs := make([]error, len(validatorErrs))
	for x, e := range validatorErrs {
		translatedErrs[x] = stderr.NewSpecific(e.Field(), e.Translate(*i.trans))
	}
	return stderr.NewMultipleError(translatedErrs...)
}

func (i *impl) InstallDefault() {
}
