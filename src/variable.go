package main

/*
// 부호 없는 정수
uint8
uint16
uint32
uint64
uint => 머신에 맞는 비트로 할당

// 부호 있는 정수
int8
int16
int32
int64
int => 머신에 맞는 비트로 할당

// 부동소수점
float32
float64

// 복소수, 32+32, 64+64
complex64
complex128

// uint와 같은 크기를 갖는 포인터
uintptr

// T/F
bool

// Unicode, 32bit
rune

// uint8  과 동일
byte

// 문자열, 수정 할 수 없음.
string

*/
import "fmt"

func main_2() {
	// 변수 선언
	var integer int
	var float float32

	// 값 할당
	integer = 3
	float = 5.5

	println(integer)
	println(float)
	fmt.Println(float)

	var a, b, c int
	fmt.Println(a, b, c)

	// := 사용 시 var 를 빼야 함.
	d, e, f := 3, "Hello", true
	fmt.Println(d, e, f)

	z := 3 + 1i // 1i 인 경우 1을 반드시 붙여야 됨.
	fmt.Println(z)

	const (
		aa = "aa"
		bb = 'B' // 아스키코드 값을 저장?
		cc = false
	)

	fmt.Println(aa, bb, cc)

	const (
		aaa = iota // enum 처럼 사용 가능
		bbb
		ccc
	)

	fmt.Println(aaa, bbb, ccc)

	// 아래와 같이 변수 선언은 불가
	// var zzz, xxx, cccc int, bool, string

	str := "HELLO world"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(bytes, str2)

	// 포인터
	var ptr = &integer
	*ptr += 7
	fmt.Println(integer)
}