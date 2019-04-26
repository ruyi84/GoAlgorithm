package main

func reverse(x int) int {
	nums := 0
	for {
		nums = (x % 10) + nums*10
		x = x / 10
		if x == 0 {
			break
		}
	}

	if nums > 2147483647 || nums < -2147483648 {
		return 0
	}
	return nums

}

func main() {
	i := 1231
	reverse(i)
}
