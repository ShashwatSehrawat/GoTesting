package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	// sums := []int{}
	var sums []int
	if len(numbersToSum) == 0 {
		sums = append(sums, 0)
	}
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
