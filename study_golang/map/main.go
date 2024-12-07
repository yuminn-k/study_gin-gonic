package main

import (
	"fmt"
)

func main() {
	// map은 키-값 쌍을 저장하는 해시테이블 자료구조입니다.
	// 선언 방법: map[key타입]value타입
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	// 다른 map 선언 방법들:
	// var m2 map[string]int					// nil map - 읽기만 가능, 쓰기 불가
	// m3 := make(map[string]int)			// 빈 map - 읽기, 쓰기 가능
	// m4 := make(map[string]int, 5)	// 초기 용량이 5인 map 생성

	// 기본적인 map 조작
	fmt.Println(m)             // 전체 map 출력
	fmt.Println(m["James"])    // 특정 키의 값 접근
	fmt.Println(m["Barnabas"]) // 존재하지 않는 키 접근 시 해당 타입의 zero value 반환

	// 안전한 map 접근 방법
	// ok는 키가 존재하는지 여부를 나타내는 bool 값
	v, ok := m["Barnabas"]
	fmt.Println(v, ok) // 0 false

	// map 조회 시 권장되는 패턴
	if v, ok := m["Miss MoneyPenney"]; ok {
		fmt.Println("Found:", v) // Found: 27
	}

	// map에 새로운 항목 추가
	m["todd"] = 33

	// map 순회
	// 참고: map은 순서가 없기 때문에 순회 순서는 보장되지 않습니다.
	for k, v := range m {
		fmt.Printf("키: %v, 값: %v\n", k, v)
	}

	// map에서 항목 삭제
	delete(m, "James")

	// 존재하지 않는 키 삭제 시 아무 일도 일어나지 않고 에러도 발생하지 않습니다.
	delete(m, "Ian Fleming")
	
	// map의 길이 확인
	fmt.Printf("map의 길이: %v\n", len(m))

	// 모범 사례: map이 nil인지 확인
	var nilMap map[string]int
	if nilMap == nil {
		fmt.Println("map이 초기화되지 않았습니다")
	}

	// 모범 사례: map 클리어하기
	for k := range m {
		delete(m, k)
	}
	
	// 또는 새로운 map으로 재할당
	// m = make(map[string]int)
}
