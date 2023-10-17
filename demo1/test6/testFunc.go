package main

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("welcome", "hello")
	println(a, b)
	c, _ := swap("ok", "are you")
	println(c)
}
