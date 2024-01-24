package main

import (
	"fmt"
)

const (
	a = iota // a는 0: iota가 0에서 시작
	b        // b는 1: 이전 iota 값에 1을 더함
	c        // c는 2: 이전 iota 값에 1을 더함
)

const (
	d = iota // d는 0: 새로운 const 블록에서 iota가 다시 0에서 시작
	e        // e는 1: 이전 iota 값에 1을 더함
	f        // f는 2: 이전 iota 값에 1을 더함
)

func main() {
	fmt.Println(a)                      // 0
	fmt.Println(b)                      // 1
	fmt.Println(c)                      // 2
	fmt.Printf("%T\n%T\n%T\n", a, b, c) // int, int, int: 모든 iota 값은 int 타입
	fmt.Println(d)                      // 0
	fmt.Println(e)                      // 1
	fmt.Println(f)                      // 2
}
