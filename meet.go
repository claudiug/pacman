package main

import (
	"fmt"
	"math"
	"image"
)

type Vertex struct {
	Y    float64
	X    float64
	name string
}

func (v Vertex) Abs() float64 {
	return math.Sqrt((float64)(v.X*v.X + v.Y*v.Y))
}

func (v *Vertex) Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := &Vertex{10, 20, "c"}
	v.Y = 100
	v.X = 200
	v.Scale(10)
	fmt.Println(v.X, v.Y)

	img := image.NewRGBA(image.Rect(0,0,100, 100))
	fmt.Println(img.Bounds())
	fmt.Println(img.At(0,0).RGBA())

}
