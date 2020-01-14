package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
executeReadFile()
fmt.Println("Nunca me voy a imprimir")

}

func executeReadFile() {
	ejecucion := readFile()
	fmt.Println(ejecucion)
}

func readFile() bool {
	archivo, err := os.Open("./holaa.txt")
	//defer archivo.Close()

	defer func() {
		archivo.Close()
		fmt.Println("Defer")

		r := recover()
		fmt.Println(r)

	}()


	if err != nil{
		panic(err)
	}

	scanner := bufio.NewScanner(archivo)
	var i int
	for scanner.Scan(){
		i++
		linea := scanner.Text()
		fmt.Println(i)
		fmt.Println(linea)
	}
	return true
	//archivo.Close()
	///return true
}