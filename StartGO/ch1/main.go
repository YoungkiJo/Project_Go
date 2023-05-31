package main

import "fmt"

/* ---------------------------------------------------------- */
// Lecture 10 struct

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	// p1 := person{"Nico", 19, favFood}
	p1 := person{name: "Nico", age: 19, favFood: favFood}
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}

/* ---------------------------------------------------------- */
// Lecture 9 map

// func main() {
// 	nico := map[string]string{"name": "nico", "age": "12"}
// 	for _, value := range nico {
// 		fmt.Println(value)
// 	}
// 	fmt.Println(nico["name"])
// }

/* ---------------------------------------------------------- */
// Lecture 8 Array, Slice
// 고정된 길이의 배열: Array / 고정되지 않은 길이의 배열: Slice

// func main() {
// 	names := [5]string{"nico", "lynn", "dal"}
// 	fmt.Println(names)
// 	names[3] = "alala"
// 	fmt.Println(names)

// 	slice_names := []string{"nico", "lynn", "dal"}
// 	fmt.Println(slice_names)
// 	slice_names = append(slice_names, "3", "4")
// 	fmt.Println(slice_names)
// }

/* ---------------------------------------------------------- */
// Lecture 7(포인터) 본 내용은 응용에서 많이 사용할 수 있으므로 코드 X

/* ---------------------------------------------------------- */
// Lecture 6
// func canIDrink(age int) bool {
// 	// if age < 18 {
// 	// 	return false
// 	// } else {
// 	// 	return true
// 	// }

// 	// if koreanAge := age + 2; koreanAge < 18 { // variable expression
// 	// 	return false
// 	// } else {
// 	// 	return true
// 	// }

// 	switch age {
// 	case 10:
// 		return false
// 	case 18:
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	fmt.Println(canIDrink(18))
// }

/* ---------------------------------------------------------- */
// Lecture 5

// func supperAdd(numbers ...int) int {
// 	var sum int = 0
// 	// for _, number := range numbers {
// 	// 	sum += (number)
// 	// }
// 	for i := 0; i < len(numbers); i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

// func main() {
// 	total := supperAdd(1, 2, 3, 4, 5, 6)
// 	fmt.Println(total)
// }

/* ---------------------------------------------------------- */
// Lecture 4

// defer 실험
// func lenAndUpper(name string) {
// 	defer fmt.Println("I'm done") // 함수의 실행이 끝난 후 실행되는 선언
// 	length := len(name)
// 	uppercase := strings.ToUpper(name)

// 	fmt.Println(length)
// 	fmt.Println(uppercase)
// }

// // naked return
// // func lenAndUpper(name string) (length int, uppercase string) {
// // 	defer fmt.Println("I'm done")
// // 	length = len(name)
// // 	uppercase = strings.ToUpper(name)

// // 	return
// // }

// func main() {
// 	// totalLength, up := lenAndUpper("nico")
// 	// fmt.Println(totalLength)
// 	// fmt.Println(up)

// 	lenAndUpper("nico")
// }

/* ---------------------------------------------------------- */
// Lecture 0 ~ 3

// func multiply(a, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// func main() {
// 	// 함수 call 시 대문자 사용(public 처럼 사용하기)
// 	fmt.Println("Hello World")

// 	something.SayHello()
// 	something.SayBye()

// 	// 상수 변수 차이 표시
// 	const name1 string = "nico" // 상수
// 	var name2 string = "nico"   // 변수
// 	name3 := "nico"             // 변수의 다른 표현
// 	fmt.Println(name1)
// 	fmt.Println(name2)
// 	fmt.Println(name3)

// 	fmt.Println(multiply(2, 2))

// 	totalLength, upperName := lenAndUpper("nico")
// 	fmt.Println(totalLength, upperName)
// }
