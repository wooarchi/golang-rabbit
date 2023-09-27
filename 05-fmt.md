# fmt 패키지를 이용한 텍스트 입출력
Go는 fmt 패키지를 사용해 간편하게 표준 입출력을 처리할 수 있다.

## fmt 패키지
- 패키지 호출
```
import "fmt"
```
- fmt 패키지 함수
  - Print(): 입력값 출력
  - Println(): 입력값 출력 및 개행
  - Printf(): 서식에 맞도록 입력값 출력
```
package main

import "fmt"

func main() {
  var a int = 10
  var b int = 20
  var f float64 = 32799438743.8297

  fmt.Print("a:", a, "b:", b)
  fmt.Println("a:", a, "b:", b "f:", f)
  fmt.Printf("a: %d b: %d f: %f\n", a, b, f)

아웃풋
a:10b:20a: 10 b: 20 f: 3.279943874382973+10
a:  10 b: 220 f:32799438743.829700
```

## 서식문자
- Printf( 서식 문자열, 인수1, 인수2, ...)
- 주요 서식문자열
  - %d
  - %f
  - %s
  - %v (기본 서식에 맞춰줌.)
 
  - 
