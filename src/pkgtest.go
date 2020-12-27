package main

// 여러개의 package 를 한번에 import 할 때는
// 소괄호를 사용한다.
import (
	"mypkg"
	"fmt"

	// 패키지 로드 시 해당 패키지의 init 만 호출 하고 싶은 경우
	// _ alias 를 붙인다.
	// 이 경우 export 된 심볼 사용이 불가능 하다
	// _ "mypkg"

	// 패키지 마지막 이름이 같은 경우
	// alias 지정이 가능하다
	// _ "other/xlib"
	// mongo "other/mongo/db"
	// mysql "other/mysql/db"
)

func main() {
	fmt.Println(mypkg.Getstring("C"))
}