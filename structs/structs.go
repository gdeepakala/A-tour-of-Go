package main

import "fmt"

type Vertex struct {
	x float32
	y float32
}

func (v Vertex) toInt32() (int32, int32) {
	return int32(v.x), int32(v.y)

}

func (p *Vertex) Scale(f float32) {
	p.x *= f
	p.y *= f
}

func main() {
	v := Vertex{1.24, 2.36}
	fmt.Println(v)
	fmt.Println("Struct accessed using dot operator:", v.x, v.y)
	ptr := &v
	fmt.Println(*ptr)
	fmt.Println("Struct accessed using pointer:", ptr.x, ptr.y)
	res1, res2 := v.toInt32()
	fmt.Println("Int32 values:", res1, res2)
	ptr.Scale(2.0)
	fmt.Println("After Scale(2.0)", ptr.x, ptr.y)
	fmt.Println("Struct Slice:")
	structSlice := []Vertex{
		{1.2, 2.3},
		{3.4, 4.5},
		{5.6, 6.7},
	}
	fmt.Println(structSlice)
	for i := 0; i < len(structSlice); i++ {
		fmt.Println(structSlice[i].x, "\t", structSlice[i].y)
	}

}
