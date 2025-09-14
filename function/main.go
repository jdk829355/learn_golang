package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// go 함수 정의 기본 문법
func valueOfPi(multiplier uint) float32 {
	// 명시적 타입 변환 (uint -> float32)
	return 3.14159 * float32(multiplier)
}

// 여러 개의 값을 반환하는 함수
func nameAndAge(uid int) (string, int) {
	switch uid {
	case 0:
		return "정대균", 23
	case 1:
		return "신짱구", 5
	default:
		return "", -1
	}
}

// 함수에 함수 전달하기
func runMath(a int, b int, op func(int, int) int) int { return op(a, b) }
func div(a int, b int) int                            { return a / b }

// defer 키워드
// 함수 반환 직전에 실행하는 키워드
// 주로 자원 해제 시 사용
func printZenOfPython() {
	filepath := "./test.txt"
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("파일을 여는 중 오류가 발생했습니다: %v", err)
	}
	// 파일 자원 해제
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 제네릭 프로그래밍
// 컴파일 시점에 타입 T가 정해지며, 에러를 방지하기 위해 T의 종류에 제한을 걸 필요가 있음

// 타입 제한 선언
type Integer interface {
	int8 | int16 | int32 | int64 | int | uint
}

func add[T Integer](a, b T) T {
	return a + b
}

func main() {
	value := valueOfPi(2)
	fmt.Println(value)

	fmt.Println("==================================================")

	// 왼쪽에 새로운 변수가 하나 이상 존재하면 무조건 := 사용
	myName, myAge := nameAndAge(0)
	fmt.Println("name:", myName)
	fmt.Println("age:", myAge)

	fmt.Println("==================================================")

	// 둘 다 선언된 변수면 =기호 사용
	var shinName string
	var shinAge int
	shinName, shinAge = nameAndAge(1)
	fmt.Println("name:", shinName)
	fmt.Println("age:", shinAge)

	fmt.Println("==================================================")

	// 받고싶지 않은 반환값은 _ 처리
	_, age := nameAndAge(12)
	if age == -1 {
		fmt.Println("등록되지 않은 유저입니다.")
	}

	fmt.Println("==================================================")

	// 함수에 함수 넘겨주기
	a, b := 1, 2
	result := runMath(a, b, div)
	fmt.Printf("a / b = %d\n", result)

	fmt.Println("==================================================")

	// defer 키워드 사용 함수 호출
	printZenOfPython()

	fmt.Println("==================================================")

	// 제네릭 사용
	fmt.Println(add(a, b))
	fmt.Println(add(uint(a), uint(b)))
}
