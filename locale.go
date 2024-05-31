package utils

type Foo struct {
	verbosity int
}

type option func(*Foo) interface{}


type Locale interface {
	MustLocalize(opts ...OptionLocale) (res string)
	GetLang(lang string) string
}
