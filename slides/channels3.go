package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
func main() {
	foo := hablar("Foo") // HL
	bar := hablar("Bar") // HL
	for i := 0; i < 5; i++ {
		fmt.Println(<-foo)
		fmt.Println(<-bar)
	}
	fmt.Println("FIN")
}
// END1 OMIT

// START2 OMIT
func hablar(msg string) <-chan string { // Devuelve un canal de strings que solo recibe. // HL
	c := make(chan string)
	go func() { // Lanzamos la rutina dentro de la funcion. // HL
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // HL
}
// END2 OMIT

