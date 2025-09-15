package main

import (
	"fmt"
)

// 이렇게 한번에 묶어서 실행 가능
var (
	// 정적 전역 변수를 초기화와 함께 선언한 경우
	age = 23

	// 정적 전역 변수를 초기화 없이 선언한 경우
	name string
)

// 상수 선언 문법
// Go에서 변수는 선언 시 기본 값으로 초기화되므로, 상수는 초기화하지 않고 선언하는 것이 불가하다.
const (
	job = "개발자"
)

func main() {
	// age는 초기화 되어있으므로 바로 사용 가능
	fmt.Println("my age is", age)
	age = 24
	fmt.Println("actually, my age is", age)

	// name은 초기화 하고 사용
	// 초기화하지 않으면 타입별 기본 값으로 초기화됨 (string은 "")
	name = "정대균"
	fmt.Println("my name is", name)

	// 상수는 변경 불가
	fmt.Println("my job is", job)

	// 지역변수 (코드 영역이 지나면(블록, 함수) 삭제됨)
	// 해당 변수에 대한 참조도 삭제됨
	// 위에 서술된 전역변수 문법 사용 가능
	// 추가로 := 연산자 사용
	myGpa := 4.3
	fmt.Println("my gpa is", myGpa)

	// golang은 같은 타입의 값을 연속적인 메모리 공간에 저장하는 방식으로 배열을 구성함
	// 하지만 개발자의 편의를 위해 내부 배열을 품고있는 slice라는 자료형을 제공함
	// []T과 같은 형식으로 나타내며 길이를 따로 명시하지 않음
	todo := []string{"예소 작업하기", "golang 스터디 준비하기", "쉬기"}
	todo = append(todo, "IT 트렌드 읽기")
	fmt.Println(todo)

	// 슬라이스의 내부 배열이 다 차면 메모리 공간을 확장하는 시도를 함
	// 하지만 사용할 공간이 없다면 다른 공간에 슬라이스를 통째로 copy를 함
	// 그래서 길이가 명확한 경우 make함수를 이용하여 미리 용량과 길이를 명시하는 것도 방법
	// 세번째 인자를 명시하지 않으면 길이와 용량이 동일한 슬라이스가 생성됨
	var a = make([]string, 10, 100)
	// zero 값으로 초기화됨
	fmt.Println(a)
	// 길이를 10으로 설정했기 때문에 11번째 요소에 접근하면 에러를 일으킴
	//fmt.Println(a[10]) // -> panic: runtime error: index out of range [10] with length 10
}
