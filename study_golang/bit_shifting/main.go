package main

import (
	"fmt"
)

const (
	_  = iota             // iota의 초기값 0을 사용하지 않고 무시합니다.
	kb = 1 << (iota * 10) // 1번째 iota 값(1)에 10을 곱한 값(10)만큼 1을 왼쪽으로 비트 시프트하여 킬로바이트(kb) 값을 설정합니다. 2^10 = 1024
	mb = 1 << (iota * 10) // 2번째 iota 값(2)에 10을 곱한 값(20)만큼 1을 왼쪽으로 비트 시프트하여 메가바이트(mb) 값을 설정합니다. 2^20 = 1048576
	gb = 1 << (iota * 10) // 3번째 iota 값(3)에 10을 곱한 값(30)만큼 1을 왼쪽으로 비트 시프트하여 기가바이트(gb) 값을 설정합니다. 2^30 = 1073741824
)

func main() {
	x := 2
	// 2진수
	fmt.Printf("%d\t\t%b", x, x) // 2		10

	var number uint = 1 // 0001 in binary

	// 왼쪽으로 2비트 시프트
	leftShifted := number << 2 // 0100 in binary, or 4 in decimal

	// 오른쪽으로 1비트 시프트
	rightShifted := number >> 1 // 0000 in binary, or 0 in decimal

	fmt.Printf("원래의 수: %d, 왼쪽으로 2비트 시프트: %d, 오른쪽으로 1비트 시프트: %d\n", number, leftShifted, rightShifted) // 원래의 수: 1, 왼쪽으로 2비트 시프트: 4, 오른쪽으로 1비트 시프트: 0

	kb1 := 1024
	mb1 := 1024 * kb
	gb1 := 1024 * mb

	// 0이 10개씩 증가
	fmt.Printf("%d\t\t\t%b\n", kb1, kb1) // 1024                    10000000000
	fmt.Printf("%d\t\t\t%b\n", mb1, mb1) // 1048576                 100000000000000000000
	fmt.Printf("%d\t\t%b\n", gb1, gb1)   // 1073741824              1000000000000000000000000000000
}
