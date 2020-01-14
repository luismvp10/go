package main

import (
	"fmt"
)

func main ()  {

	//matriz := [] int {1,2,3,4}
	//
	//if matriz == nil {
	//	fmt.Println("Esta vacio")
	//}else{
	//	fmt.Println(len(matriz))
	//}
	//fmt.Println(matriz)


	arreglo := [3] int {1,2,3}
	slice := arreglo[:2]
	fmt.Println(slice)

}
