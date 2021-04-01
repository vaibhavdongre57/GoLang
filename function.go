package main

import "fmt"

type Point1 struct {
	X float64
	Y float64
}
// receiver argument
func (p Point1) IsAbove(y float64)bool{
	return  p.Y > y
}
/*
  Translates the current Point, at location (X,Y), by dx along the x axis and dy along the y axis
  so that it now represents the point (X+dx,Y+dy).
*/
func (p *Point1) Translate(dx , dy float64){   // Method with Pointer receiver
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// Function with Pointer argument
func TranslateFunc(p *Point1, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func main(){
	//p := Point1{2.0,4.0}
	//
	//fmt.Println("Point: ",p)
	//
	//fmt.Println("Is Point p located above the line y = 1.0 ? : ",p.IsAbove(1))

	//p := Point1{3, 4}
	//fmt.Println("Point p: ",p)
	//
	//p.Translate(7, 6)
	//fmt.Println("After Translate, p: ", p)

	//Methods with Pointer receivers vs Functions with Pointer arguments
	p := Point1{3,4}
	ptr := &p

	fmt.Println("Point p: ", p)

	// Calling a Method with Pointer receiver
	p.Translate(2,3)
	ptr.Translate(5,10)

	// Calling a Function with a Pointer argument
	TranslateFunc(ptr, 20, 30)
	//TranslateFunc(p, 1, 2)//not valid p


}
