package main

import "fmt"

func main() {
	cad := "arepera"
	fmt.Println(cad)
	fmt.Println("******")
	fmt.Println(len(cad))
	fmt.Println("******")
	fmt.Println(cad[0:5])
	fmt.Println("******")
	fmt.Println("Ingrese un numero")
	var name int
	fmt.Scanln(&name)
	fmt.Println("******")
	fmt.Println(name)
}