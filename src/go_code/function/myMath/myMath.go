package myMath

func Add(number ...int) (sum int) {
	sum = 0
	for _, num := range number {
		sum = sum + num
	}
	return
}
