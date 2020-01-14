package main

import (
	"fmt"
	"strconv"
	)

func main() {
	edad := "22"
	edad_int, _ := strconv.Atoi(edad)
	//edad :=22
	//edad_str := strconv.Itoa(edad)
	//nombre := "Coco"
	//fmt.Println("Mi edad es :" + edad_str + " Mi nombre es : "+ nombre )
	fmt.Println(edad_int + 10)

}