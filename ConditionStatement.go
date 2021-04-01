package main

import "fmt"

//if condition
//func main(){
//	var num = 25
//	if(num%5 == 0){
//		fmt.Printf("%d is multiple of 5 \n",num)
//	}
//
//	var age =25
//	if age >=17 && age<=25{
//		fmt.Println("my age is between 17 and 30")
//	}
//
//}

//if else condition
//func main()  {
//	var age = 18
//	if age >= 18 {
//		fmt.Print("You are eligible for voting")
//	} else{
//		fmt.Print("You are not eligible for voting")
//	}
//
//}

//if-else-if chain
//func main(){
//	var WT = 22.5
//	if WT < 18.5{
//		fmt.Print("You are underweight")
//	}else if WT >= 18.5 && WT < 25.0{
//		fmt.Print("Your weight is normal")
//	}else if WT >= 25.0 && WT <30.0{
//		fmt.Print("You are overweight")
//	}else{
//		fmt.Print("You are obese")
//	}
//}

//If with a short statement
//func main(){
//	if a := 10; a%2 == 0 {
//		fmt.Printf("%d is even no.",a)
//	}
//}

//func main(){
//	if a:=15; a%2 ==0{
//		fmt.Printf("%d is even no.",a)
//	}else{
//		fmt.Printf("%d is odd no.",a)
//	}
//}

//Switch case
//func main(){
//	var dayofweek = 8
//	switch dayofweek{
//	case 1: fmt.Print("Monday")
//	case 2: fmt.Print("tuesday")
//	case 3: fmt.Print("wednesday")
//	case 4: fmt.Print("thrusday")
//	case 5: fmt.Print("friday")
//	case 6: {
//		fmt.Println("Its Saturday")
//		fmt.Println("yaaay....")
//	}
//	case 7: {
//		fmt.Println("wohho Sunday")
//		fmt.Println("yaaay....")
//	}
//	default:
//		fmt.Println("invalid day")
//	}
//}

//Switch with a short statement
//func main(){
//	switch dayofweek := 5; dayofweek{
//	case 1: fmt.Print("Monday")
//	case 2: fmt.Print("tuesday")
//	case 3: fmt.Print("wednesday")
//	case 4: fmt.Print("thrusday")
//	case 5: fmt.Print("friday")
//	case 6: {
//		fmt.Println("Its Saturday")
//		fmt.Println("yaaay....")
//	}
//	case 7: {
//		fmt.Println("wohho Sunday")
//		fmt.Println("yaaay....")
//	}
//	default:
//		fmt.Println("invalid day")
//	}
//}

//Combining multiple Switch cases
//func main(){
//	switch dayofweek := 6; dayofweek {
//	case 1, 2, 3, 4, 5: fmt.Print("WEEK-Day")
//	case 6, 7:
//		fmt.Print("WEEK-End")
//	default:
//		fmt.Print("Invalid")
//	}
//}

//Switch with no expression
//func main(){
//	var wt = 8
//	switch {
//	case wt < 18:
//		fmt.Print("Wou are under weight")
//	case wt >= 18 && wt < 25:
//		fmt.Print("Your weight is normal")
//	case wt >25 && wt <	30:
//		fmt.Print("you are over weighted")
//	default:
//		fmt.Print("you are obese")
//
//	}
//}


//For loop
//func main(){
//	for i:=1; i<=10; i++{
//		fmt.Printf("%d\t",i)
//	}
//}

//Omitting the initialization statement
//func main(){
//	i := 2
//	for ;i <= 10; i+=2{
//		fmt.Printf("%d\t",i)
//	}
//}

//Omitting the increment statement
//func main(){
//	i := 2
//	for ;i<=10;{
//		fmt.Printf("%d\t",i)
//		i *= 2
//	}
//	fmt.Print("\n 'i' is greater that 10")
//}

//func main() {
//	i := 2
//	for i <= 20 { //  ';' is removed from condition
//		fmt.Printf("%d ", i)
//		i *= 2
//	}
//}


//Break statement
//func main(){
//	for num:=1; num <=100; num++{
//		if num %3 ==0 && num%5 == 0{
//			fmt.Printf("First positive number divisible by 3 and 5: %d",num)
//			break
//			print("'if' loop is breaked...")
//		}
//	}
//}


//Continue Statement
func main(){
	for i:=1; i <= 10; i++{
		if i%2 == 0{
			fmt.Print("aaya kya")
			break
			fmt.Printf("continue loop")
		}
		fmt.Printf("%d\t",i)
	}

}