package main

// typedef 같은 개념
// typedef int func(int, int)
type calc func(int, int) int

// calc 와 같은 유형의 함수를 인자로 받을 수 있음
func wrapper(f calc, v1 int, v2 int) int {
	return f(v1, v2)
}

// int f(void) 형식의 함수를 반환
func nextValue() func() int {
	var i int = 0

	// closure
	return func() int {
		i++ // c++ 람다와 달리 외부의 모든 걸 캡쳐하는 듯.
		return i
	}
}

func main_6() {
	// 람다 같은 개념
	sum := func(n ...int) int {
		s := 0
		for _, value := range n {
			s += value
		}

		return s
	}

	println(sum(1, 2, 3, 4, 5))
	println(wrapper(func(x int, y int) int {return x- y}, 10, 20))

	next := nextValue()
	for i := 0; i < 3; i++ {
		println(next())
	}

	next = nextValue()
	for i := 0; i < 3; i++ {
		println(next())
	}
}