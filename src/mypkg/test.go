package mypkg

// package 단위는
// GOPATH 의 src 밑에 디렉터라 단위로 한다.

// 소문자로 시작하는 변수나 함수는 외부에서 접근 할 수 없다.
// 외부로 export 할 심볼은 대문자로 시작해야 한다.

// init() 은 패키지 로드 시 제일 처음 호출되는 함수

var pop map[string]string

type Template struct {
	//wow string // 소문자로 시작하면 이 변수를 접근 할 수 없게 됨.
	Wow string
}

func init() {
	println("MYPKG INIT")

	pop = map[string]string {
		"A": "Hello",
		"B": "Good",
		"C": "Bye",
	}
}

func Getstring(key string) string {
	return pop[key]
}