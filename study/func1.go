package main

import "fmt"

func add1(x, y int) int {
	return x + y
}
func main() {

	fmt.Println(add1(5, 6))

	var num1 int
	var num2, num3 int
	var num4, num5, str1 = 4, 5, "hello"
	var (
		i int
		b bool
		s string = "hi"
	)

	fmt.Println("①", num1)
	fmt.Println("②", num2, num3)
	fmt.Println("③", num4, num5, str1)

	num6 := 6
	fmt.Println("④", num6)
	fmt.Println("⑤", i, b, s)
}
