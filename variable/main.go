package main

import (
	"fmt"
)

// 정적 전역 변수를 초기화와 함께 선언한 경우
var age = 23

// 정적 전역 변수를 초기화 없이 선언한 경우
var name string

func main() {
	// age는 초기화 되어있으므로 바로 사용 가능
	fmt.Println(age)
	age = 24
	fmt.Println(age)

	// name은 초기화 하고 사용
	// 초기화하지 않으면 타입별 기본 값으로 초기화됨 (string은 "")
	name = "정대균"
	fmt.Println(name)
}
