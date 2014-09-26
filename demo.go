package main

import (
	"fmt"
	"runtime"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func maina() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{10, 23}
	fmt.Println(v.X)
	fmt.Println(v.Y)

	p := Vertex{1, 1}
	q := &p
	q.X = 10
	fmt.Println(p.X)
	fmt.Println(p.X)
	vv := new(Vertex)
	fmt.Println(vv)

	var a [2]string
	fmt.Println(a)
	a[0] = "Claudiu"
	fmt.Println(a)

	ss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(ss)
	fmt.Println(ss[1:len(ss)])

	for x, y := range ss {
		fmt.Printf("%d\n - %d\n", x, y)
	}

	demo := func(name string) string {
		return strings.ToUpper(name)
	}
	fmt.Println(demo("claudiu"))

	fmt.Println(runtime.GOOS)
}
