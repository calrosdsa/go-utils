package locale

import (
	_r "github.com/calrosdsa/go-utils"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

type locale struct {
	bundle *i18n.Bundle
	defaultLanguage string
	locales []string
}


func New() _r.Locale {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	locales := viper.GetStringSlice("locales")
	for _,locale := range locales {
	    bundle.MustLoadMessageFile(fmt.Sprintf("./locale/active.%s.toml",locale))
	}
	// lo := locales[0]
	return &locale{
		bundle: bundle,
		defaultLanguage: locales[0],
		locales: locales,
	}
}

func (l *locale) GetLang(lang string) string{
	for _,locale := range l.locales {
		if locale == lang {
			return lang 
		}
	}
	return l.defaultLanguage
}

func (l *locale) MustLocalize(opts ..._r.OptionLocale) (res string) {
	options:= _r.OptionsLocale.Apply(opts...)
	if options.Lang == "" {
		options.Lang = l.defaultLanguage
	}
	localizer := i18n.NewLocalizer(l.bundle, options.Lang)
	res = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: options.ID,
			One: options.One,
			Other: options.Other,
		},
		TemplateData: options.Template,
	})
	return
}
