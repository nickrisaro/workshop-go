package main

import (
	"fmt"
	"time"
)

// START OMIT
func generaMensajes(c chan string) {
	for i := 0 ; ; i++ {
		c <- fmt.Sprintf("Mensaje generado %d", i) // HL
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan string)
	go generaMensajes(c)

	for i := 0; i < 5; i++ {
		fmt.Println("Mensaje recibido: ", <-c) // HL
	}

	fmt.Printf("Dejé de escuchar")
}
// END OMIT