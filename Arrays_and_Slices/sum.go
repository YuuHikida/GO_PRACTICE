package main

func Sum(numbers []int) int {
	sum := 0

	// "_"は配列の添え字を無視している
	for _, number := range numbers {
		sum += number
	}

	// 添え字を表示するver↓
	// for x, number := range numbers {
	// 	sum += number
	// 	fmt.Printf("(%d)", x)
	// }
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	countNum := len(numbersToSum)
	sumSlice := make([]int, countNum)

	for i, num := range numbersToSum {
		sumSlice[i] += Sum(num)
	}
	return sumSlice
}
