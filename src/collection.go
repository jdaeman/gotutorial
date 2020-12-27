package main

import "fmt"

func main_7() {

	// 배열
	// [3]int 와 [5]int 는 서로 다른 타입이다.
	// 배열의 크기를 동적으로 증가 시킬 수 없음.
	// 부분 배열을 발췌 할 수 없음.
	var a [3]int // int a[3]
	for idx, _ := range(a) {
		a[idx] = idx
		println(a[idx])
	}

	var a1 = [3]int{4, 5, 6} // 배열 길이 선언 및 값 초기화
	var a2 = [...]int{7, 8, 9} // 배열 크기가 자동으로 설정 됨

	println(a1[0])
	println(a2[1])

	var adim2 = [2][3]int {
		{1, 2, 3},
		{4, 5, 6}, // C++과 달리 끝이어도 , 를 붙임
	}

	// 2차원 배열 순회
	for _, arr := range adim2 {
		for _, value := range arr {
			print(value, " ")
		}
		println()
	}

	// 슬라이스
	// 배열의 제약점을 뛰어넘음
	// 고정된 크기를 미리 지정하지 않음
	// 이후 그 크기가 동적으로 변하고, 
	// 부분 배열을 발췌 할 수 있다.
	var sl []int // 길이가 명시 되어 있지 않음
	if sl == nil {
		println("nil ", len(sl), " ", cap(sl))
	}
	
	sl = []int{10, 11, 12, 13}
	println(sl[3])

	// slice, length, capacity
	ssl := make([]int, 5, 10)
	println(len(ssl), " ", cap(ssl))
	// capacity 를 벗어나서 runtime error
	println(ssl[4]/*, ssl[9], ssl[11]*/)

	sl = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// range => [2, 5) idx 2 부터 4 까지만 
	sssl := sl[2:5]
	for idx, value := range sssl {
		println(idx, " ", value)
	}

	// [0, len(s1))
	sssl = sl[0:]
	println(sssl)
	fmt.Println(sssl)
	println()

	sliceA := make([]int, 0, 3)

	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}

	sliceB := []int{7, 7, 7}
	// slice 에 slice 를 append 할 때는 ... 이 필요하다.
	sliceA = append(sliceA, sliceB...) 
	fmt.Println(sliceA)

	source := []int{0, 1, 2}
	{
		var target []int
		copy(target, source)
		// 내부 배열의 길이와 용량이 없으므로 복사 되지 않음.
		fmt.Println(target)
	}

	target := make([]int, len(source), cap(source) * 2)
	copy(target, source)
	fmt.Println(target)
	
	// map
	// 따로 초기화 없이 선언만 하면
	// nil map 이라 해서 메모리가 할당되어 있지 않음
	// [key]value
	var m map[int]string

	m = make(map[int]string)
	m[0] = "A"
	m[1] = "B"
	m[2] = "C"

	for key, val := range m {
		fmt.Println(key, " ", val)
	}

	// key 를 이용하여 값과 존재 여부를 확인 가능
	_, exist := m[5]
	fmt.Println(exist)

	// key 를 이용하여 값 확인
	// 없는 경우 value type 의 nil 을 반환.
	val := m[2]
	fmt.Println(val)

	delete(m, 777)
	delete(m, 1) // key 가 있는 경우 삭제함
	for key, val := range m {
		fmt.Println(key, " ", val)
	}

	m = map[int]string {
		1: "A",
		2: "B",
		3: "C",
	}
	fmt.Println(m)
}