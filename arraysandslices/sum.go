package arraysandslices

func Sum(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value
	}
	return
}

func SumAll(numberSlices ...[]int) (sums []int) {
	for _, numbers := range numberSlices {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numberSlices ...[]int) (sums []int) {
	for _, numbers := range numberSlices {
		if len(numbers) < 1 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return
}
