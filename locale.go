package utils

import "net/http"

type Foo struct {
	verbosity int
}

type option func(*Foo) interface{}

type Locale interface {
	MustLocalize(opts ...OptionLocale) (res string)
	GetLang(r *http.Request) string
}
