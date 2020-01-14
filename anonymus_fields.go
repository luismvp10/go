package main

import (
	"fmt"
)

type Human struct {
	name string
}

func (this Human) hablar() string  {
	return "Bla bla blah"
}


type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (this Tutor) hablar() string  {
	return "Bienvenido: "+ this.Human.hablar()
}


func main()  {

	tutor := Tutor{Human{"Uriel"}, Dummy{"Luis"}}

	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.hablar())
}
