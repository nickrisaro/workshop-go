package main

import "fmt"
// START OMIT
func main() {

	var s []int
	fmt.Println(s)

	if s == nil {
		fmt.Println("nil!")
	}

	a := make([]int, 0) // Crea un slice vacío
	fmt.Println(a)

	b := make([]int, 5) // Crea un slice con 5 enteros (zero values)  
	fmt.Println(b)
}
// END OMIT
