package main

import (
	"fmt"
	"awesomeProject/example/model"
)
func main() {
	c := model.Customer{
		Id: 1,
		Name: "Rajeev Singh",
	}

	c.Married = true	// Error: can not refer to unexported field or method

	//a := model.address{} // Error: can not refer to unexported name

	fmt.Println("Programmer = ", c);
}
