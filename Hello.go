package utils

import "fmt"

type Hello struct {
	Name string
}

func HelloPerson(name string) {
	fmt.Println(name)
}

func HelloPerson2(name string){
	fmt.Sprintf("Hello %s",name)
}