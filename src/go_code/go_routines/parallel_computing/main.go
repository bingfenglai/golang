package main

// 数列求和函数
func sum(list []int, ch chan int) error {
	i := 0

	for _, j := range list {
		i = i + j
	}

	ch <- i
	return nil
}

// 参与运算的cpu数
const cpu_num = 6

func main() {

	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 18}

}
