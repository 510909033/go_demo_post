package my_benchmark

//计算数的和
func Add(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
