package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["hoge"] = Vertex{
		10, 10.0,
	}

	fmt.Println(m)
}
