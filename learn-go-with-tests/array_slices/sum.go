package array_slices

func Sum1(numbers [5]int) int {

	sum := 0

	for i := 0; i < 5; i++ {

		sum += numbers[i]
	}

	return sum
}

func Sum2(numbers [5]int) int {

	sum := 0

	for _, number := range numbers {

		sum += number
	}

	return sum
}

func Sum3(numbers []int) int {

	sum := 0

	for _, number := range numbers {

		sum += number
	}

	return sum
}

func SumAll(numbersSum ...[]int) []int {

	lengthnumbersSum := len(numbersSum)
	sums := make([]int, lengthnumbersSum)

	for i, number := range numbersSum {

		sums[i] = Sum3(number)
	}

	return sums
}

func SumAll1(numbersSum ...[]int) []int {

	var sums []int
	for _, number := range numbersSum {

		sums = append(sums, Sum3(number))
	}
	return sums
}

func SumAllTails(numbersSum ...[]int) []int {

	var sums []int
	for _, number := range numbersSum {

		sums = append(sums, Sum3(number[1:]))
	}
	return sums
}

func SumAllTails1(numbersSum ...[]int) []int {

	var sums []int
	for _, number := range numbersSum {

		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum3(number[1:]))
		}

	}
	return sums
}
