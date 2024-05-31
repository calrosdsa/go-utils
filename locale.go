package utils

type Foo struct {
	verbosity int
}

type option func(*Foo) interface{}


type Locale interface {
	MustLocalize(id string,lang string) (res string)
}
