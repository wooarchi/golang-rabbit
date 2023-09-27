package main

import "fmt"

func main() {
    // 문자열 출력
    fmt.Println("Fmt Test")

    // 변수 값 출력
    name := "John"
    age := 30
    fmt.Printf("이름: %s, 나이: %d\n", name, age)

    // 포맷된 문자열 생성
    formattedString := fmt.Sprintf("이름: %s, 나이: %d", name, age)
    fmt.Println(formattedString)

    // 서식 지정자를 사용한 출력
    floatNumber := 3.14159
    fmt.Printf("원주율: %.2f\n", floatNumber) // 소수점 둘째 자리까지 출력

    // 값을 가진 변수 출력
    fmt.Printf("이름: %v, 나이: %v\n", name, age)

    // 타입 정보와 값 출력
    fmt.Printf("이름의 타입: %T, 나이의 타입: %T\n", name, age)

    // 개행 없이 출력
    fmt.Print("이 문자열은 개행 없이 출력됩니다.")
    fmt.Print("이 문자열도 개행 없이 출력됩니다.")
}
