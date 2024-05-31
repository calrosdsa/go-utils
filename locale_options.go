package utils

type optionsLocale struct {
	Lang     string
	ID       string
	One      string
	Other    string
	Template interface{}
}

type OptionLocale func(o *optionsLocale)
var OptionsLocale optionsLocale

func (o *optionsLocale) WithLang(lang string) OptionLocale {
	return func(o *optionsLocale) {
		o.Lang = lang
	}
}

func (o *optionsLocale) WithID(id string) OptionLocale {
	return func(o *optionsLocale) {
		o.ID = id
	}
}

func (o *optionsLocale) WithOne(one string) OptionLocale {
	return func(o *optionsLocale) {
		o.One = one
	}
}

func (o *optionsLocale) WithOther(other string) OptionLocale {
	return func(o *optionsLocale) {
		o.Other = other
	}
}

func (o *optionsLocale) WithTemplate(template interface{}) OptionLocale {
	return func(o *optionsLocale) {
		o.Template = template
	}
}

func (o *optionsLocale) Apply(opts ...OptionLocale) optionsLocale {
	options := optionsLocale{}
	for _,opt := range opts {
		opt(&options)
	}
	return options
}