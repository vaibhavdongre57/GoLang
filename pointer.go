package main

import "fmt"

func main(){

	//var p *int
	//fmt.Println("P=",p)


	// Initializing a Pointer
	var a = 100
	var p = &a

	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("Value stored in variable p =", p)
	//Dereferencing a Pointer
	fmt.Println("Value stored in variable *p =", *p)

	fmt.Println("a (before) = ", a)

	// Changing the value stored in the pointed variable through the pointer
	*p = 2000

	fmt.Println("a (after) = ", a)


	//Creating a Pointer using the built-in new() function
	ptr := new(int)
	*ptr = 1234

	fmt.Printf("Ptr=%#x, Pointer value = %d\n",ptr,*ptr)


	var a1 = 7.98
	var p1 = &a1
	var pp = &p1

	fmt.Println("a = ", a1)
	fmt.Println("address of a1 = ", &a1)

	fmt.Println("p1 = ", p1)
	fmt.Println("address of p1 = ", &p1)

	fmt.Println("pp = ", pp)

	// Dereferencing a pointer to pointer
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)//** is used to return value in ptr to ptr

	//var x1 = 67
	//var p2 = &x1
	//
	////var q1 = p2 + 1 // Compiler Error: invalid operation
	//println(q1)

	var z = 75
	var v1 = &z
	var v2 = &z

	if v1 == v2 {
		fmt.Println("Both pointers p1 and p2 point to the same variable.")
	}
}