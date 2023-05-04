package main

import "fmt"

func main() {
	sampleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	numberFilter := getNewFilter(sampleInput)

	// even
	fmt.Println(numberFilter.isEven().filterAll())

	// odd
	// fmt.Println(numberFilter.isOdd().filterAll())

	// even prime
	// fmt.Println(numberFilter.isEven().isPrime().filterAll())

	// odd prime
	// fmt.Println(numberFilter.isOdd().isPrime().filterAll())

	// multiple
	// fmt.Println(numberFilter.isMultipleOf(2).filterAll())

	// lesser
	// fmt.Println(numberFilter.isLesserThan(10).filterAll())

	// greater
	// fmt.Println(numberFilter.isGreaterThan(10).filterAll())

	//custom
	// fmt.Println(numberFilter.custom(func(number int) bool {
	// 	return number > 5 && number < 10
	// }).filterAll())

	// ALL
	// fmt.Println(numberFilter.isGreaterThan(10).isPrime().filterAll())

	// Any
	// fmt.Println(numberFilter.isGreaterThan(10).isPrime().filterAny())

}
