package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {

	fmt.Println("Future investment program")

	var investmentAmount, years, fv, rfv float64
	// var expectedReturnRate = 5.5
	expectedReturnRate := 5.5 //shorthand for inferred variables

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter Years: ")
	fmt.Scan(&years)

	fv, rfv = calcFutureValues(investmentAmount, expectedReturnRate, years)

	formattedVal := fmt.Sprintf("Future Value: %.2f\nFuture Adjusted Value %.2f\n", fv, rfv)
	outputText(formattedVal)
	// fmt.Printf("Future Value: %.2f\n Future Adjusted Value %.2f\n", futureValue, futureAdjustedValue) outputText(formattedVal)

}

func outputText(formattedVal string) {

	fmt.Print(formattedVal)
}

func calcFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {

	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return

}
