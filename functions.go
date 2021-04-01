package main

import (
	"errors"
	"fmt"
	"math"
)

func avg(a, b float64)float64{
	return (a+b)/2
}

//Functions with multiple return values
func getStockPriceChange(prevPrice, currentPrice float64)(float64,float64){
	change := currentPrice - prevPrice
	percentChange := (change/prevPrice) * 100
	return change, percentChange
}

//Returning an error value from a function
func getStockPriceChangeWithError(prevPrice, currentPrice float64)(float64,float64, error){
	if prevPrice == 0{
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice-prevPrice
	precentChange := (change / prevPrice) * 100
 	return change, precentChange, nil
}

// Function with named return values
func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice //Notice how we changed := (short declarations) with = (assignments) in the function body
	percentChange = (change / prevPrice) * 100
	return change, percentChange
}

func main(){
	//println(avg(22.5,28.65))
	//a := 22.5
	//b := 28.65
	//result := avg(a, b)
	//fmt.Printf("Average of %.2f and %.2f is %.2f\n",a,b,result)

	prevStockPrice := 80000.00
	currentStockPrice := 120000.00

	//Functions with multiple return values
	//change, precentChange := getStockPriceChange(prevStockPrice, currentStockPrice)
	//if change < 0{
	//	fmt.Printf("the stock price decreased by %.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(precentChange))
	//}else{
	//	fmt.Printf("the stock price increased by %.2f which is %.2f%% of the prev price\n",change, precentChange)
	//}


	//Returning an error value from a function
	//change, percentChange, err := getStockPriceChangeWithError(prevStockPrice, currentStockPrice)
	//if err != nil{
	//	fmt.Println("Sorry there was an error", err)
	//}else{
	//	if change < 0{
	//		fmt.Printf("the stock price decreased by %.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	//	}else{
	//		fmt.Printf("the stock price increased by %.2f which is %.2f%% of the prev price\n",change, percentChange)
	//	}
	//}

	//Functions with named return values
	//change, percentChange := getNamedStockPriceChange(prevStockPrice, currentStockPrice)
	//
	//if change < 0 {
	//	fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	//} else {
	//	fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	//}

	//Blank Identifier
	//Sometimes you may want to ignore some of the results from a function that returns multiple values
	//But in that case, you’ll be forced to use the percentChange variable because Go doesn’t allow creating variables that you never use.
	//So what’s the solution? Well, you can use a blank identifier
	change, _ := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f\n", math.Abs(change))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f\n", change)
	}
}

