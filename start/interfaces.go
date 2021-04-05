package main

import ("fmt")

type Shape interface {
	x int
}



func main() {
	shape  := Shape{1, 2}

	fmt.Println(shape.x)
}