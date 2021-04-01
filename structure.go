package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
	Age int
}

type Car struct {
	Name, Model, Color string
	WeightInKg float64
}

type Student struct {
	RollNo int
	Name string
}

type Employee struct {
	Id int
	Name string
}

type Point struct {
	X float64
	Y float64
}

func main() {

	// Declaring a variable of a `struct` type
	var p Person // // All the struct fields are initialized with their zero value
	fmt.Println(p)

	// Declaring and initializing a struct using a struct literal
	var p1 = Person{"Vaibhav","Dongre",23}
	fmt.Println("Person1: ",p1)

	// Naming fields while initializing a struct
	p2 := Person{
		FirstName: "Tejaswini",
		LastName: "Dongre",
		Age: 26,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)

	c := Car{
		Name: "Baleno",
		Model: "Alfa",
		Color: "White",
		WeightInKg: 1254,
	}

	//Struct Car
	// Accessing struct fields using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value to a struct field
	c.Color = "Black"
	fmt.Println("Car: ", c)

	//Struct Student
	// instance of student struct type
	s := Student{63,"Vaibhav"}
	fmt.Println("Student1: ", s)

	// Pointer to the student struct
	ps := &s
	fmt.Println(ps)

	// Accessing struct fields via pointer
	fmt.Println((*ps).Name)
	fmt.Println(ps.Name) // Same as above: No need to explicitly dereference the pointer

	ps.RollNo = 31
	fmt.Println(ps)

	//You can also get a pointer to a struct using the built-in new() function
	// It allocates enough memory to fit a value of the given struct type, and returns a pointer to it
		pEmp := new(Employee)

		pEmp.Id = 1000
		pEmp.Name = "Sachin"

		fmt.Println(pEmp)

	g1 := Point{10,20}
	g2 := g1 // A copy of the struct `p1` is assigned to `p2`
 	fmt.Println("Value of g1: ",g1)
	fmt.Print("Value of g2: ",g2)

	g2.X = 15
	fmt.Println("\nAfter modifying g2")
	fmt.Println("Value of g1: ",g1)
	fmt.Print("Value of g2: ",g2)

	// Two structs are equal if all their corresponding fields are equal.
	h1 := Point{10.50,20.50}
	h2 := Point{10.50,20.50}

	if h1==h2{
		fmt.Println("\nPoint p1 and p2 are equal.")
	} else {
		fmt.Println("Point p1 and p2 are not equal.")
	}

}