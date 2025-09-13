package main

import (
	"fmt"
)

func main() {
	name := "정대균"
	a := make([]string, 10, 100)

	// conditional statement

	// if문 -> 그냥 if문
	if a[0] == "" {
		fmt.Println("slice a's first value is zero value")
	} else {
		fmt.Println("slice a's first value is not zero value")
	}
	// switch문 -> c와 달리 break 명시하지 않아도 됨, default 필수
	switch name {
	case "정대균":
		fmt.Println("Hello! 정대균")
	case "홍길동":
		fmt.Println("Hello! 홍길동")
	default:
		fmt.Println("Who are you??")
	}
}
