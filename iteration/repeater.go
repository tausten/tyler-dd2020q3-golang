package iteration

func Repeat(toRepeat string, timesParams ...int) (repeated string) {
	times := 5
	if len(timesParams) > 0 {
		times = timesParams[0]
	}
	for i := 0; i < times; i++ {
		repeated += toRepeat
	}

	return
}
