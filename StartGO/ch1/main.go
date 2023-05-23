package main

import (
	"fmt"
	"githubcode/StartGO/ch1/something"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	// 함수 call 시 대문자 사용(public 처럼 사용하기)
	fmt.Println("Hello World")

	something.SayHello()
	something.SayBye()

	// 상수 변수 차이 표시
	const name1 string = "nico" // 상수
	var name2 string = "nico"   // 변수
	name3 := "nico"             // 변수의 다른 표현
	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("nico")
	fmt.Println(totalLength, upperName)

}
