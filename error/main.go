package main

import (
	"errors"
	"fmt"
)

// errors 모듈로 error 값 반환
// 에러 변수 이름은 Err___로 하는 것이 권장됨
// 안 지킬 시 오류는 뜨지 않지만 노란 줄이 거슬림
var ErrDivisionByZero = errors.New("division by zero")
var ErrorSample = errors.New("Hello")

// 0으로 나눴을 때를 보여주기 위한 예시 함수
func Devide(number, d float32) (float32, error) {
	if d == 0 {
		// 0으로 나눈 경우 에러 반환
		return 0, ErrDivisionByZero
	} else {
		// 정상적인 경우 error의 default value인 nil 반환
		return number / d, nil
	}
}

// 예상치 못한 에러가 발생했을 때는 panic을 발생시킨다.
// panic이 발생하면 함수를 즉시 반환하며 최초 호출 지점까지 연달아 반환된다.
// main함수까지 멈추는 경우를 막기 위하여 (서버를 어떻게든 살리기 위하여) recover함수를 쓴다.

// 여기서 발생한 패닉은 recover되지 않고 즉시 반환되어 testRecover로 전달된다.
func testPanic() {
	panic(ErrorSample)
	fmt.Println("Hello from testPanic")
}

// recover함수는 defer키워드를 통해 사용된다.
// panic이 발생해도 defer키워드에 있는 함수는 실행된다.
// 이를 이용하여 panic이 main까지 전파되는 것을 막을 수 있다.
func testRecover() {
	defer func() {
		fmt.Printf("recover(): %v\n", recover())
		// if recover() != nil {
		// 	fmt.Println("got an error")
		// } else {
		// 	fmt.Println("no error")
		// }
	}()
	testPanic()
	fmt.Println("Hello from testRecover")
}

func main() {
	n1, e1 := Devide(1, 1)
	fmt.Println("result:", n1)
	if e1 != nil {
		fmt.Println(e1.Error())
	}
	n2, e2 := Devide(1, 0)
	fmt.Println("result:", n2)
	if e2 != nil {
		// error인터페이스의 Error()함수를 호출하여 에러 메시지 출력
		fmt.Println(e2.Error())
	}
	testRecover()
}
