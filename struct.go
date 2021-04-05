package main

import ("fmt")

type Point struct {
	x int32
	y int32
	isTrue bool
}


type Circle struct  {
	point *Point
}


func main() {

	// var point = Point{1, 2, false}
	// p1 := Point{x:1}
	c1 := Circle{point: &Point{x:1}}
	fmt.Println(c1.point)
}