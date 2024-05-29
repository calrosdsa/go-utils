package hello

import "fmt"

type Hello struct {
	Name string
}

func HelloPerson(name string) {
	fmt.Println(name)
}