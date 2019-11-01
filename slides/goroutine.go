package main

import (
	"fmt"
	"time"
)

func say(s *string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(*s)
	}
}

func main() {
	mensaje := "mundo"
	go say(&mensaje)

	mensaje2 := "hola"
	say(&mensaje2)
}
