package main

import "fmt"

type Vertice struct {
	Lat, Long float64
}

var m = map[string]Vertice{
	"Bell Labs": Vertice{
		40.68433, -74.39967,
	},
	"Google": Vertice{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
