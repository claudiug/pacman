package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main(){
	fmt.Println(Vertex{1, 2})
	v := Vertex{10,23}
	fmt.Println(v.X)
	fmt.Println(v.Y)

	p := Vertex{1,1}
	q := &p
	q.X = 10
	fmt.Println(p.X)
	fmt.Println(p.X)
	vv := new(Vertex)
	fmt.Println(vv)
}

