package main

import "fmt"

func main() {
    // 문자열 변수 선언 및 초기화
    var greeting string
    greeting = "안녕하세요, Go 언어!"

    // 정수 변수 선언 및 초기화
    var number1 int
    number1 = 42

    // 변수 선언과 동시에 초기화
    var number2 int = 23
    var isTrue bool = true

    // 자료형 추론을 사용하여 변수 선언 및 초기화
    name := "John"
    age := 30

    // 변수 값 출력
    fmt.Println(greeting)
    fmt.Println("number1:", number1)
    fmt.Println("number2:", number2)
    fmt.Println("isTrue:", isTrue)
    fmt.Println("name:", name)
    fmt.Println("age:", age)

    // 변수 값 변경
    number1 = 55
    fmt.Println("변경된 number1:", number1)

    // 상수 선언
    const pi = 3.14159
    fmt.Println("원주율:", pi)
}
