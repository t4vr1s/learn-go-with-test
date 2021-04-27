package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	for _, number := range numbersToSum {
		if len(number) == 0 {
			sum = append(sum, 0)
		} else {
			tail := number[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}
