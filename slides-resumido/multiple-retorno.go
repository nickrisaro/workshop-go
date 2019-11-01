package main

import "fmt"

func intercambiar(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := intercambiar("Grupo", "Esfera")
	fmt.Println(a, b)
}

