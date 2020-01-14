package main

import (
	"fmt"
)

func main()  {

	/*
		1. Es una dirección de memoria
		2. En lugar del valor, tenemos la dirección en la que está el valor
		3. X, Y => #as123d =>
		4. X => #as123 => 6
		5. Y ¿? => 6
		*T => *int, *string, *Struct
		valor zero => nil
	*/

	var x,y *int
	entero := 5

	x= &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)
}
