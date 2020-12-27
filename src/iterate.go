package main

/*
* go lang 에서 반복문은 for 만 존재
*/

func main_4() {
	var sum int = 0
	
	/*
	for (int i = 1; i <= 100; i++) {
		sum += i
	}
	*/
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	n := 1
	for n < 100 {
		n *= 2
		if n > 60 {
			break
		} else {
			println("Continue")
			continue
		}
	}
	println(n)

INFINITE:
	for {
		// 무한 루프
		// 반복문 에 정의 된 label 에 대한 break 시
		// 해당 반복문을 탈출
		// goto 를 이용해서 탈출 할 수 도 있음
		// 이 경우는 아래 쪽에 있는 label 을 명시해야 함.
		break INFINITE
	}
	
	// 배열
	// range keyword 를 이용해서 배열을 순회 할 수 있음
	hw := []string{"cpu", "memory", "disk"}
	for index, name := range hw {
		println(index, name)
	}

	for _, name := range hw {
		println(name)
	}

	for index, _ := range hw {
		println(index)
	}

	for index := 0; index < 3; index++ {
		println(hw[index])	
	}
}