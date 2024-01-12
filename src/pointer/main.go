package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)					// 42
	fmt.Println(&a)					// 메모리 주소. 0x14000124008
	fmt.Printf("%T\n", a)		// int
	fmt.Printf("%T\n", &a)	// *int

	// var b int = &a	// 컴파일 에러. cannot use &a (type *int) as type int in assignment
	var b *int = &a	// *int 타입의 포인터 변수 b에 a의 메모리 주소를 할당
	fmt.Println(b)	// 메모리 주소. 0x14000124008
	fmt.Println(*b)	// 메모리 주소가 가리키는 값. 42
	fmt.Println(*&a)	// 메모리 주소가 가리키는 값. 42
}