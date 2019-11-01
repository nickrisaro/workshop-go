package main

import "fmt"

func intercambiar(x, y string) (primera string, segunda string) {
	primera = y
	segunda = x
	return
}

func main() {
	a, b := intercambiar("mercado", "libre")
	fmt.Println(a, b)
}

