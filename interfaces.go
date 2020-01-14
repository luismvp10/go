package main

import (
	"fmt"
)

type User interface {
	Permisos() int   //1-5
	Nombre() string
}

type Admin struct {
	nombre string
}

type Editor struct {
	nombre string
}


func (this Editor) Permisos() int  {
	return 3
}

func (this Editor) Nombre() string {
	return this.nombre
}




func (this Admin) Permisos() int  {
	return 5
}

func (this Admin) Nombre() string {
	return this.nombre
}

func auth(user User) string  {
	permisos :=user.Permisos()

	if permisos > 4 {
		return user.Nombre() + " tiene permisos de administrador"
	}else if permisos == 3 {
		return user.Nombre() + " tiene permisos de editor"
	}
	return ""
}

func main()  {

	//admin := Admin{"Uriel"}
	//editor := Editor{"Fulano"}

	usuarios := [] User{Admin{"Uriel"}, Editor{"Fulano"}}

	for _,usuario := range usuarios{
		fmt.Println(auth(usuario))
	}

	//fmt.Println(auth(admin))
	//fmt.Println(auth(editor))
}