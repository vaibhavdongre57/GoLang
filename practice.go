package main

import "fmt"

/*
func main(){
	fmt.Println("hello")
}
*/
//
//func main(){
//	var num1 int = 10
//	num2 := 20
//	fmt.Println(num1)
//	fmt.Println(num2)
//	fmt.Println(num1+num2)
//}


//data type
//func main(){
//	var(
//		fname, lname  string = "Vaibhav","Dongre"
//		age			  int = 23
//		salary        float64 = 80000
//		isconfirmed   bool = true
//	)
//
//	fmt.Printf("First name: %s\n Last name= %s\n age= %d\n Salary= %f\n isConfirmed= %t\n",fname,lname,age,salary,isconfirmed)
//
//}


////Type inference
//func main(){
//	var name="vaibhav"
//	//name= 212
//	fmt.Printf("Variable 'name' is type of %T\n",name)
//}
//
//func main() {
//	var fname, lname, age, salary = "vaibhav", "dongre", 23, 80000.0
//
//	fmt.Printf("Fname:%T, Lname:%T, age=%T, salary=%T",fname, lname, age, salary)
//}


//Short Declaration
func main(){
	name := "vaibhav" // Short variable declaration syntax ':='
	age, salary, isConfirmed := 23, 8000.0, true

	fmt.Println(name, age, salary, isConfirmed)
}