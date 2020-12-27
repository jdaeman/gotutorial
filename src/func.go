package main

import "fmt"

// call by value
func cbv(msg string) {
	fmt.Println(msg)
}

// call by reference 
func cbr(msg * string) {
	*msg += *msg
}

// 가변인자
// 동일한 타입에 대해서만
func variadic(msg ...string) {
	for index, str := range msg {
		println(index, str)
	}
}

// mulitple return
func adder(nums ...int) (int, int) {
	s := 0
	count := 0

	for _, num := range nums {
		count++
		s += num
	}

	return count, s
}

// Named Return Parameter
func adder_v2(nums ...int) (count int, total int) {
	for _, num := range nums {
		count += 1
		total += num
	}
	// 이 경우 반드시 return 을 작성해주어야 한다.
	return
}

func main() {
	str := "Hello"

	cbr(&str)
	cbv(str)

	variadic("This", "is", "a", "book")

	{
		count, total := adder(1, 2, 3, 4, 5)
		println(count, total)
	}

	_, total := adder_v2(6, 7, 8, 9)
	println(total)
}