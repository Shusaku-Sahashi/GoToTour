package main

import "fmt"

type vertex struct{
	Lat, Long float64
}

m = map[string]vertex {
	"google": Vertex{
		40, 20.0,
	},
	"hoge": Vertex {
		10, 10.0
	},
}

func main() {
	fmt.Println(m)
}