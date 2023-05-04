package main

import (
	"reflect"
	"testing"
)

type numberFilterTest struct {
	output         []int
	description    string
	expectedOutput []int
}

func TestNumberFilter(t *testing.T) {
	sampleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	testCases := []numberFilterTest{
		{
			output:         getNewFilter(sampleInput).isEven().filterAll(),
			description:    "testing isEven",
			expectedOutput: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
		{
			output:         getNewFilter(sampleInput).isOdd().filterAll(),
			description:    "testing isOdd",
			expectedOutput: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
		},
		{
			output:         getNewFilter(sampleInput).isPrime().filterAll(),
			description:    "testing isPrime",
			expectedOutput: []int{2, 3, 5, 7, 11, 13, 17, 19},
		},
		{
			output:         getNewFilter(sampleInput).isGreaterThan(10).filterAll(),
			description:    "testing isGreaterThan",
			expectedOutput: []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		},
		{
			output:         getNewFilter(sampleInput).isLesserThan(10).filterAll(),
			description:    "testing isLesserThan",
			expectedOutput: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			output:         getNewFilter(sampleInput).isMultipleOf(5).filterAll(),
			description:    "testing isMultipleOf",
			expectedOutput: []int{5, 10, 15, 20},
		},
		{
			output:         getNewFilter(sampleInput).isOdd().isPrime().filterAll(),
			description:    "testing isOdd AND isPrime",
			expectedOutput: []int{3, 5, 7, 11, 13, 17, 19},
		},
		{
			output: getNewFilter(sampleInput).custom(func(number int) bool {
				return number > 1 && number < 6
			}).filterAny(),
			description:    "testing custom",
			expectedOutput: []int{2, 3, 4, 5},
		},
		{
			output:         getNewFilter(sampleInput).isOdd().isPrime().filterAny(),
			description:    "testing isOdd OR isPrime",
			expectedOutput: []int{1, 2, 3, 5, 7, 9, 11, 13, 15, 17, 19},
		},
		{
			output:         getNewFilter(sampleInput).isOdd().isPrime().isGreaterThan(10).filterAll(),
			description:    "testing isOdd AND isPrime AND isGreaterThan",
			expectedOutput: []int{11, 13, 17, 19},
		},
	}

	for _, testCase := range testCases {

		if !reflect.DeepEqual(testCase.output, testCase.expectedOutput) {
			t.Fatal("Test Failed: " + testCase.description)
		}
	}

}
