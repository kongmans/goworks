package main

func add(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
func main() {
	a := add(1, 2, 3, 4, 5,6)
	println(a)
}
