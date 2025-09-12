// main 패키지의 파일이라는 뜻
package main

// 모듈 임포트 부분
// fmt: format이라는 뜻으로  i/o 기능을 담당함
import (
	"fmt"
)

// 파일의 진입점인 main 함수
func main() {
	// Println() 함수, fmt 모듈에서 가져온 함수이며 인자를 출력하고 개행하는 기능을 가짐.
	// func Println(a ...any) (n int, err error)
	// 인자간 " "을 추가하고 return value로 출력의 바이트 수와 에러를 가짐
	fmt.Println("Hello, World!")
}
