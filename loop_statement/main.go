package main

import (
	"fmt"
)

func main() {
	names := []string{"정대균", "홍길동", "김철수"}

	// := range 방법
	// python의 enumerate와 같은 문법
	// 메서드 호출이나 별도 카운터 변수가 필요없어서 효율적임
	for i, v := range names {
		fmt.Println("index: ", i)
		fmt.Println("value: ", v)
	}

	fmt.Println("====================================")

	// 전통적인 방법도 가능
	for i := 0; i < len(names); i++ {
		fmt.Println("index: ", i)
		fmt.Println("value: ", names[i])
	}

	fmt.Println("====================================")

	i := 0
	// while문은 따로 없고 다음과 같이 사용
	for i < len(names) {
		fmt.Println("index: ", i)
		fmt.Println("value: ", names[i])
		i++
	}
}
