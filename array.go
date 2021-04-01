package main
import "fmt"

func main() {
	var x [5]int // An array of 5 integers
	fmt.Println(x)

	var y [8]string // An array of 8 strings
	fmt.Println(y)

	var z [3]complex128 // An array of 3 complex numbers
	fmt.Println(z)

	//Accessing array elements by their index
	 var a [5] int
	a[0] = 100
	a[1] = 101
	a[3] = 103
	a[4] = 104
	fmt.Printf("a[0]= %d,  a[1] = %d,  x[2] = %d\n",a[0],a[1],a[2])
	fmt.Println("a = ",a)

	// Declaring and initializing an array at the same time
	var d = [5]int{2, 4, 6, 8, 10}
	fmt.Println(d)

	// Short declaration for declaring and initializing an array
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println(b)

	// You don't need to initialize all the elements of the array.
	// The un-initialized elements will be assigned the zero value of the corresponding array type
	c := [5]int{2}
	fmt.Println(c)

	// Letting Go compiler infer the length of the array
	e := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(e)

	//Arrays in Golang are value types
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)

	//Iterating over an array in Golang
	names := [3]string{"Mark Zuckerberg", "Bill Gates", "Larry Page"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// len() function
	f := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for i := 0; i < len(f); i++ {
		sum = sum + f[i]
	}
	fmt.Printf("Sum of all the elements in array  %v = %f\n", f, sum)

	//Iterating over an array using range
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek{
	//for _, value := range daysOfWeek {  //another way of writing for loop
		fmt.Printf("Day %d of week = %s\n", index, value)
	}


	//Multidimensional arrays in Golang
	g := [2][2]int{
		{3, 5},
		{7, 9},	// This trailing comma is mandatory
	}
	fmt.Println(g)

	// Just like 1D arrays, you don't need to initialize all the elements in a multi-dimensional array.
	// Un-initialized array elements will be assigned the zero value of the array type.
	h := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(h)



}
