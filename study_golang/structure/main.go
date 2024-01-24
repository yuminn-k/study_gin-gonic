package main

import (
	"fmt"
)

// person 구조체는 사람의 이름과 나이를 나타냅니다.
type person struct {
	first string // 이름
	last  string // 성
	age   int    // 나이
}

// secretAgent 구조체는 person 구조체를 임베딩하여, 사람의 기본 정보 외에
// 비밀 요원 여부를 나타내는 ltk 필드를 추가합니다.
type secretAgent struct {
	person // person 구조체를 임베딩하여 상속과 유사한 효과를 냅니다.
	first  string
	ltk    bool // 비밀 요원 여부 (License to Kill)
}

func main() {
	// p1은 person 구조체의 인스턴스로, 이름과 성, 나이 정보를 가집니다.
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	// p2는 person 구조체의 인스턴스로, 필드 이름을 생략하고 값을 순서대로 초기화합니다.
	p2 := person{"Miss", "Moneypenny", 27} // 비추천. 필드 이름을 생략하면 가독성이 떨어집니다.

	// p1과 p2의 정보를 출력합니다.
	fmt.Println(p1) // {James Bond 32}
	fmt.Println(p2) // {Miss Moneypenny 27}

	// p1의 각 필드를 개별적으로 출력합니다.
	fmt.Println(p1.first, p1.last, p1.age) // James Bond 32

	// sa1은 secretAgent 구조체의 인스턴스로, person 정보와 ltk 값을 가집니다.
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "something coll",
		ltk:   true, // 비밀 요원 여부를 true로 설정합니다.
	}

	// sa1의 person 필드와 ltk 값을 출력합니다.
	// 임베딩된 구조체의 필드에 직접 접근이 가능합니다.
	fmt.Println(sa1.person.first, sa1.first, sa1.last, sa1.age, sa1.ltk) // James something coll Bond 32 true

	// Anonymous struct
	// 구조체를 정의하지 않고, 필요한 필드만을 가지는 구조체를 생성할 수 있습니다.
	anon := struct {
		name string
		age  int
	}{
		name: "Jane Doe",
		age:  30,
	}

	// 익명 구조체의 필드를 출력합니다.
	fmt.Println(anon.name, anon.age) // Jane Doe 30
}
