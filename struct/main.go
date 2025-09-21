package main

import (
	"fmt"
)

// 기본적인 구조체 정의 방법
type User struct {
	Name string
	Age  int
}

// 사용자가 정의 하여 타입처럼 사용 가능
func nameAndAge(uid uint) User {
	switch uid {
	case 0:
		// 사용 시에는 중괄호로 감싸서 사용
		return User{"정대균", 23}
	case 1:
		return User{"신짱구", 5}
	default:
		return User{"", -1}
	}
}

// go의 구조체는 값타입이므로 인자로 구조체를 받을 때 복사본을 전달받음
// func incrementAge(user User) {
// 	// 여기서 user는 main의 user의 복사본임
// 	user.Age++
// 	fmt.Println("incrementAge() 호출", user.Age)
// }
/*
[출력]
호출 전
내 이름은 신짱구, 5살이죠
incrementAge() 호출 6
호출 후
내 이름은 신짱구, 5살이죠
*/

// User의 포인터인 *User를 인자로 받음
func incrementAge(user *User) {
	// 역참조를 따로 안 함 (구조체의 포인터 한정)
	// 컴파일러가 알아서 user.Age -> (*user).Age로 변환해주기 때문에 따로 역참조를 할 필요 없음
	// 이는 syntactic sugar(구문 설탕)의 한 종류
	user.Age++
	fmt.Println("incrementAge() 호출", user.Age)
}

/*
[출력]
호출 전
내 이름은 신짱구, 5살이죠
incrementAge() 호출 6
호출 후
내 이름은 신짱구, 6살이죠
*/

// 이 경우에는 int의 주소를 출력함
func printNum(i *int) {
	fmt.Println(i)
}

// 값 리시버와 포인터 리시버

// 값 리시버는 구조체의 복사본을 전달받아 처리함
// 그러므로 값 조작 불가능
func (user User) prettyString() {
	fmt.Printf("안녕, 내 이름은 %s이고 나이는 %d세야. \n", user.Name, user.Age)
}

// 포인터 리시버는 구조체의 포인터를 전달받아 처리함
// 그러므로 값을 조작할 수 있음
func (user *User) updateName(newName string) {
	user.Name = newName
}

func (user *User) incrementAge() {
	// 역참조를 따로 안 함 (구조체의 포인터 한정)
	// 컴파일러가 알아서 user.Age -> (*user).Age로 변환해주기 때문에 따로 역참조를 할 필요 없음
	// 이는 syntactic sugar(구문 설탕)의 한 종류
	user.Age++
	fmt.Println("incrementAge() 호출", user.Age)
}

// 인터페이스 정의
// 함수 이름, 입력 파라미터 타입, 출력 타입 정의
type Living interface {
	incrementAge()
	prettyString()
	updateName(string)
}

// 해당 함수가 문제 없이 실행되면 User타입이 인터페이스 구현을 잘 했다고 볼 수 있음
func incrementAgeAndPrettyString(being Living) {
	being.incrementAge()
	being.prettyString()
}

func main() {
	user := nameAndAge(0)
	// . 으로 요소에 접근할 수 있음

	fmt.Println("호출 전")
	fmt.Printf("내 이름은 %s, %d살이죠\n", user.Name, user.Age)

	// 구조체의 pointer는 stringer 인터페이스의 String()함수를 구현할 때 주소가 아닌 값을 보여주도록 설계되어있음
	// int같은 타입의 포인터를 출력할 때는 주소가 그대로 나옴
	p := &user
	fmt.Println(p)
	n := 100
	fmt.Println(&n)

	//user의 포인터를 전달하여 함수에서 수정할 수 있도록 함
	incrementAge(p)

	fmt.Println("호출 후")
	fmt.Printf("설날이 지났군요!! 내 이름은 %s, %d살이죠\n", user.Name, user.Age)

	// 포인터를 넘길 시 역참조가 필요 없는 경우는 구조체인 경우에만 해당됨
	printNum(&n) // 0x1400009a038

	// 값 리시버 함수 호출
	user.prettyString()

	// 포인터 리시버 함수 호출
	user.updateName("정대만")
	user.prettyString()

	// Living 인터페이스를 구현하는 User 구조체

	// 책에서는 다 값이나 포인터로 통일할 필요가 있다고 했지만 아님
	// 실제고 incrementAge와 prettyString은 각각 포인터, 값 리시버인데도 포인터를 인자로 넣으니 동작이 됨

	// 만약 &user대신 user를 사용하면 에러가 남
	// 값타입은 값타입을 받는 리시버만 사용 가능
	// 포인터 타입은 값타입과 포인터타입의 리시버 둘 다 사용 가능
	incrementAgeAndPrettyString(&user)
}
