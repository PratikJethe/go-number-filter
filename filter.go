package main

import (
	"math"
)

type numberFilter struct {
	numbers    []int
	conditions []filterFunction
}
type filterFunction func(int) bool

func (nf *numberFilter) isEven() *numberFilter {

	nf.conditions = append(nf.conditions, func(number int) bool {
		return number%2 == 0
	})

	return nf
}
func (nf *numberFilter) isOdd() *numberFilter {

	nf.conditions = append(nf.conditions, func(number int) bool {
		return number%2 != 0
	})
	return nf
}
func (nf *numberFilter) isPrime() *numberFilter {

	nf.conditions = append(nf.conditions, func(number int) bool {
		if number < 2 {

			return false
		}
		sq_root := int(math.Sqrt(float64(number)))
		for i := 2; i <= sq_root; i++ {
			if number%i == 0 {
				return false
			}
		}
		return true
	})
	return nf
}
func (nf *numberFilter) isMultipleOf(divisor int) *numberFilter {

	nf.conditions = append(nf.conditions, func(number int) bool {
		return number%divisor == 0
	})
	return nf
}
func (nf *numberFilter) isGreaterThan(num int) *numberFilter {

	nf.conditions = append(nf.conditions, func(number int) bool {
		return number > num
	})
	return nf
}
func (nf *numberFilter) isLesserThan(num int) *numberFilter {

	nf.conditions = append(nf.conditions, func(number int) bool {
		return number < num
	})
	return nf
}
func (nf *numberFilter) custom(customFilterFunction filterFunction) *numberFilter {

	nf.conditions = append(nf.conditions, customFilterFunction)

	return nf
}

func (nf *numberFilter) filterAll() []int {

	result := []int{}
numberloop:
	for _, number := range nf.numbers {

		for _, filter := range nf.conditions {

			if !filter(number) {
				continue numberloop
			}

		}
		result = append(result, number)
	}
	return result
}
func (nf *numberFilter) filterAny() []int {

	result := []int{}
	for _, number := range nf.numbers {

		for _, filter := range nf.conditions {
			// if any match happens, append number
			if filter(number) {
				result = append(result, number)
				break
			}

		}
	}
	return result
}

func getNewFilter(numbers []int) *numberFilter {

	return &numberFilter{
		numbers: numbers,
	}
}
