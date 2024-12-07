package main

import (
	"fmt"
)

// Person 구조체 정의
// 구조체는 서로 다른 타입의 데이터를 하나로 묶어서 관리할 수 있게 해줍니다
type Person struct {
	name    string  // 이름
	age     int     // 나이
	address string  // 주소
}

// Employee 구조체는 Person을 포함(임베딩)합니다
// Go에서는 상속 대신 구조체 임베딩을 통해 코드 재사용을 합니다
type Employee struct {
	Person          // Person 구조체를 임베딩
	salary  int     // 급여
	company string  // 회사명
}

func main() {
	// 1. 구조체 인스턴스 생성 방법 1 - 필드명 명시
	p1 := Person{
		name:    "김철수",
		age:     25,
		address: "서울시",
	}

	// 2. 구조체 인스턴스 생성 방법 2 - 필드 순서대로 값 입력
	// 이 방법은 가독성이 떨어져서 추천하지 않습니다
	p2 := Person{"홍길동", 30, "부산시"}

	// 3. 구조체 포인터 생성
	p3 := &Person{
		name:    "이영희",
		age:     28,
		address: "인천시",
	}

	// 4. 임베디드 구조체 생성
	emp := Employee{
		Person: Person{
			name:    "박지민",
			age:     35,
			address: "대전시",
		},
		salary:  5000000,
		company: "테크컴퍼니",
	}

	// 구조체 필드 접근
	fmt.Println("이름:", p1.name)           // 일반 구조체 필드 접근
	fmt.Println("나이:", p2.age)            // 일반 구조체 필드 접근
	fmt.Println("주소:", p3.address)        // 포인터 구조체도 동일하게 접근
	
	// 임베디드 구조체 필드 접근
	// 임베디드 필드는 직접 접근 가능합니다
	fmt.Println("직원 이름:", emp.name)      // Person의 필드 직접 접근
	fmt.Println("직원 급여:", emp.salary)    // Employee 자체 필드 접근

	// 5. 익명 구조체 생성
	// 한 번만 사용할 구조체는 익명으로 선언 가능합니다
	temp := struct {
		id   int
		data string
	}{
		id:   1,
		data: "임시 데이터",
	}
	fmt.Println("임시 데이터:", temp.data)
}
