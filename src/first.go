package main

/*
* main package 는 Go compiler 가 특별하게 인식함
* main package 는 executable program 으로 만들며
* 이 패키지 중 main() 이 프로그램의 시작점이 된다.
* 따라서 라이브러리 제작 시 에는 package main 과 main() 을 사용해서는 안된다.
*/

/*
go keyword
break
default
func
interface
select
case
defer
go
map
struct
chan
else
goto
package
switch
const
fallthrough
if
range
type
continue
for
import
return
var
*/


import "fmt"

func main_1(){
	println("Hello World")

	fmt.Println("Hello World")
}