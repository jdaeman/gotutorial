package main

import "fmt"
import "mypkg"

// typedef struct person
type person struct {
	name string
	age uint32

	// 메소드는 별도로 분리되어 정의됨.
}

/* func info() {

}
*/
// call by value
func (p person) info() {
	fmt.Println("Info name: ", p.name, " age: ", p.age)
}

// call by reference
func (p *person) ptr_info() {
	p.name += p.name
	p.info()
}

func main() {

	p := person{} // empty
	var pp person

	p.name = "Hello"
	p.age = 32

	pp = person{}

	// pointer 를 반환하지만, -> 로 접근해서 값을 변경하진 않는다.
	// struct pointer 는 일반 객체처럼 . 으로 접근한다.
	ppp := new(person)
	ppp.name = "World"

	var pppp *person
	pppp = &pp
	pppp.name = "HHH"

	fmt.Println(p)
	fmt.Println(pp)
	fmt.Println(*ppp)

	abc := mypkg.Template{}
	abc.Wow = "HHH"
	fmt.Println(abc)

	ppp.info()
	ppp.ptr_info()
}