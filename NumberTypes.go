package main

//func main(){
//	var myInt8 int8 = 97
//	var myInt = 1200
//	var uint = 500
//	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
//	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers
//
//	fmt.Printf("%d, %d, %d, %#x, %#o\n", myInt8, myInt, uint, myHexNumber, myOctalNumber)
//}

//func main(){
//	var myByte byte = 'a'
//	var myRune rune = 'A' // Type inferred as `rune` (Default type for character values)
//
//	fmt.Printf("%c = %d  and  %c = %U",myByte, myByte, myRune, myRune)
//	//printed the variable myByte in character and decimal format, and the variable myRune in character and Unicode format
//}


//operations
//func main(){
//	var a, b= 4, 5
//	var res1 = (a+b)*(a+b)/2  //Arithematic operation
//
//	a++  //increament by 1
//	fmt.Println(a)
//
//	b += 10  //increament by 10
//	fmt.Println(b)
//
//	var res2 = a^b  //Bitwise XOR
//	fmt.Println(res2)
//
//	var r =3.5
//	var res3 = math.Pi * r *r  //operations on floating point type
//	fmt.Println(math.Pi)
//
//	fmt.Printf("rest1 = %v, rest2 = %v, rest3 = %v\n",res1,res2,res3)
//
//}



//Logical Operators:
//func main(){
//	var truth = 3 <= 5  //Boolean type
//	var flasehood = 10 != 10
//
//	var test1 = 10 > 20 && 5==5  // Second operand is not evaluated since first evaluates to false
//	var test2 = 2*2==4 || 10%3 == 0  // Second operand is not evaluated since first evaluates to true
//
//	fmt.Println(truth, flasehood, test1, test2)
//}


//Complex number
//func main(){
	//var x = 5 +7i  //assume it as complex128
	//var a = 3.57
	//var b = 6.23
	//
	//// var c = a + bi won't work. Create the complex number like this -
	//var c = complex(a, b)
//
//	var a = 3 + 5i
//	var b = 2 + 4i
//
//	var res1 = (a+b)
//	var res2 = (a-b)
//	var res3 = (a*b)
//	var res4 = (a/b)
//
//	fmt.Println(res1,res2,res3,res4)
//}

//Strings
//func main() {
//	var website = "\thttps://www.callicoder.com\t\n"
//
//	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
//                           practical guides and tutorials on programming languages,
//                           web development, and desktop app development.\t\n`
//
//	fmt.Println(website, siteDescription)
//}

//Type Conversion
//func main(){
//	//var a int64 = 4
//	//var b int = a // Compiler Error (Cannot use a (type int64) as type int in assignment)
//	//var c int =500
//	//var result = a + c // Compiler Error (Invalid Operation: mismatched types int64 and int)
//
//
//	//explicitly cast the variables to the target type -
//	var a int64 = 4
//	var b int = int(a)  // Explicit Type Conversion
//
//	var c float64 = 6.5
//
//	// Explicit Type Conversion
//	var result = float64(b) + c // Works
//	fmt.Println(result)
//}


//Constant
//func main(){
//
//	const length = 10
//	const width = 20
//	var area int
//
//	area = length*width
//	fmt.Println("Value of area:",area)
//
//}
