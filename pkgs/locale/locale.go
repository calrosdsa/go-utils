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
}

func New() _r.Locale {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	locales := viper.GetStringSlice("locales")
	for _,locale := range locales {
	    bundle.MustLoadMessageFile(fmt.Sprintf("./locale/active.%s.toml",locale))
	}
	return &locale{
		bundle: bundle,
	}
}

func (l *locale) MustLocalize(id string, lang string) (res string) {
	localizer := i18n.NewLocalizer(l.bundle, lang)
	res = localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: id,
		},
	})
	return
}
