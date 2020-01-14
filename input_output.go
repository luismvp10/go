package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	var edad int
	var nombre string
	fmt.Println("Coloca tu nombre:")
	fmt.Scanf("%s\n", &nombre)
	fmt.Println("Coloca tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("Mi nombre es: %s\n", nombre)
	fmt.Printf("Mi edad es %d\n",edad)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	text,err := reader.ReadString('\n')

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+text)
	}
}
