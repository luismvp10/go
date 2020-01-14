package main

import (
	"fmt"
)


type User struct {
	edad int
	nombre, apellido string

}

func (this User) nombre_completo () string {
	return this.nombre+ " " + this.apellido
}

func (this *User) set_name(n string)  {
	this.nombre = n
}

func main()  {

	//var uriel User
	//uriel := User{nombre:"Uriel", apellido:"Sanchéz", edad:29}
	uriel := new(User)

	uriel.nombre = "Uriel"
	uriel.apellido = "Hernández"
	uriel.set_name("Marcos")

	fmt.Println(uriel.nombre)
	fmt.Println(uriel.nombre_completo())


}
