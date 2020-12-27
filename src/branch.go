package main

import "fmt"

func main_3() {
	k := 1

	// 소괄호 없이 바로 조건을 작성
	if k == 1 {
		fmt.Println("K is 1")		
	} else if k == 2 { // else if 는 반드시 이런 식으로
		fmt.Println("K is 2")
	} else {
		fmt.Println("K is other")
	}

	// if ~ else 에서만 사용 가능한 변수 초기화 가능
	if val := k << 1; val < 2 {
		fmt.Println(val, "if")
	} else {
		fmt.Println(val, "else")
	}

	// break 를 사용하지 않아도 break 가 있는 거 처럼 동작
	switch k {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("default")
	}

	// break 를 사용하지 않는 것 처럼 동작시키려면
	// fallthrough 를 사용
	switch k {
	case 1:
		fmt.Println("case 1")
		fallthrough
	case 2:
		fmt.Println("case 2")
		fallthrough
	case 3:
		fmt.Println("case 3")
		fallthrough
	default:
		fmt.Println("default")
	}

	// interface 가 아닌 변수에 대해서 type 사용 불가
	/*switch k.(type) {
	case int:

	case bool:

	case string:

	default:
	}*/

	score := 80
	// switch 뒤에 표현식을 적지 않으면
	// if ~ else if ~ else 를 간결하게 사용 할 수 있음
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	default:
		println("F")
	}
}