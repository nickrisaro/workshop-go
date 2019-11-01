package main

import "fmt"
// START OMIT
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	imprimirSlice(s)

	s = s[:0] // Corta el slice para que tenga longitud cero
	imprimirSlice(s)

	s = s[:4] // Extiende su longitud
	imprimirSlice(s)

	s = s[2:] // Remueve sus dos primeros valores
	imprimirSlice(s)
}

func imprimirSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
// END OMIT
