# Go Number Filter

Go Number Filter is an program that provides various ways to filter an given array of numbers
It uses builder pattern to nest multiple filters. program supports all or any filter type

## Features

- Filters (even, odd, prime, multiplOf, lesserThan, greaterThan, custom)
- Filters can be of ALL or ANY type
- Can nest muiltiple patterns

## How to use

- clone this repository [git clone https://github.com/PratikJethe/go-number-filter.git]
- to run program [go run .]

## Examples

### 1. isEven Filter

```sh
    sampleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numberFilter := getNewFilter(sampleInput)
	// even
	fmt.Println(numberFilter.isEven().filterAll())

	//TERMINAL
	$ go run .
    [2 4 6 8 10 12 14 16 18 20]
```

### 2. Nesting Multiple Filters

```sh
    sampleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numberFilter := getNewFilter(sampleInput)
	// even prime
	fmt.Println(numberFilter.isEven().isPrime().filterAll())

	//TERMINAL
	$ go run .
    [2]
```

### 3. Custom Filter Function

```sh
    sampleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numberFilter := getNewFilter(sampleInput)
	//custom
	fmt.Println(numberFilter.custom(func(number int) bool {
	 return number > 5 && number < 10
	}).filterAll())

	//TERMINAL
	$ go run .
    [6 7 8 9]
```

### 4. All Type Filter

```sh
    sampleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numberFilter := getNewFilter(sampleInput)
	//ALL type
	fmt.Println(numberFilter.isGreaterThan(10).isPrime().filterAll())

	//TERMINAL
	$ go run .
    [11 13 17 19]
```

### 5. Any Type Filter

```sh
    sampleInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	numberFilter := getNewFilter(sampleInput)
	//ALL type
	fmt.Println(numberFilter.isGreaterThan(10).isPrime().filterAny())

	//TERMINAL
	$ go run .
    [2 3 5 7 11 12 13 14 15 16 17 18 19 20]
```
