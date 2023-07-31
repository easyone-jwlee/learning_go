package main

import "fmt"

func main() {
	// 변수를 선언하는 방법
	var x int = 10
	fmt.Printf("--------------------: %v\n", x)

	// integer literal default type은 int이기 때문에 data type을 쓰지 않아도 int로 선언된다.
	var y = 10
	fmt.Printf("--------------------: %v\n", y)

	// 변수를 선언하고 zero value를 할당하고 싶다.
	var z int
	fmt.Printf("--------------------: %v\n", z)

	// 한번의 var line으로 같은 type으로 여러개의 변수를 선언.
	var xy, yz int = 10, 20
	fmt.Printf("--------------------:%v %v\n", xy, yz)

	// 같은 타입의 모든 변수에 zero value 할당
	var a, b int
	fmt.Printf("--------------------:%v %v\n", a, b)

	// 서로 다른 타입
	var n, s = 10, "hello"
	fmt.Printf("--------------------:%v %v\n", n, s)

	// 선언 list를 괄호로 묶어 사용
	var (
		c    int
		d        = 20
		e    int = 30
		f, g     = 40, "hello"
		h, i string
	)
	fmt.Printf("-------------------- c:%v, d:%v, e:%v, f:%v, g:%v, h:%v, i:%v\n", c, d, e, f, g, h, i)

	// := 연산자를 사용하여 type 유추를 하는 var 선언을 대신 한다.
	// var x = 10 과 x := 10 은 정확히 같은 동작을 한다.
	// x에 10의 값을 할당하며 int로 선언.

	// := 에도 한가지 제한이 있다. package level에서 변수를 선언하고자 한다면, := 연산자의 사용은 함수 밖에서는 불가능하기 때문에 반드시 var 키워드를 사용해야 한다.
}
