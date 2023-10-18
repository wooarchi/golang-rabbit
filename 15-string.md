# 문자열
- 문자열이란 문자 집합.
- 타입명 'String' 사용.
- 문자열은 백쿼트(`)로 묶어서 사용.
- 백쿼트는 내부에서는 특수문자가 일반문자로 처리된다.
```
func example() {
  str1 := "Hello\t'World'\n"

  str2 := `Go is "awesome"!\nGo is simple and\t'powerfule`

  fmt.Println(str1)
  fmt.Println(str2)
}

Hello   'World'

Go is "awesome"!\nGo is simple and\t'powerfule
```
- 큰 따옴표를 사용하여 여러 줄 문자열 출력
```
func example() {
  str := "This is a multi-line\nstring using double quotes.\nIt can include escape characters."
  fmt.Println(str)
}

This is a multi-line
string using double quotes.
It can include escape characters.
```
- 백쿼트를 사용하여 여러 줄 문자열 출력
```
func example() {
  str := `  This is a multi-line
  string using backticks.
  It is a raw string, so escape characters are not interpreted.`
  fmt.Println(str)
}

  This is a multi-line
  string using backticks.
  It is a raw string, so escape characters are not interpreted.
```
### UTF-8 문자코드
- Go는 UTF-8 문자코드를 표준 문자코드로 사용.
- UTF-8은 다국어 문자를 지원하고 문자열 크기를 절약할 목적으로 만들어진 문자코드.
- UTF-16은 한 문자에 2바이트를 고정해서 사용, UTF-8은 자주 사용되는 영문자, 숫자, 일부 특수 문자를 1바이트로 표현하고 그외 문자를 2~3바이트로 표현하여 절약.
- ANSI 코드와 1:1 대응.

### rune 타입으로 한문자 담기
- 문자 하나를 표현하는 데 rune 타입 사용.
- 4바이트 정수 타입인 int32 타입의 별칭 타입. (rune = int32)
```
func example() {
  var char rune = '한'

  fmt.Printf("%T\n", char)
  fmt.Println(char)
  fmt.Printf("%c\n", char)
}

int32
54620
한
```

### len()으로 문자열 크기 알아내기
- len()은 내장함수
- 문자열 크기는 문자의 수가 아닌 문자열이 차지하는 메모리 크기
```
func example() {
  str1 := "가나다라마" 
  str2 := "abcde"

  fmt.Printf("len(str1) = %d\n", len(str1))
  fmt.Printf("len(str2) = %d\n", len(str2))
}

len(str1) = 15 // 3byte x 5
len(str2) = 5 // 1byte x 5
```
### []rune 타입 변환으로 글자 수 알아내기 ([]rune 슬라이스 타입에 대한 이해필요.)
- string 타입, []rune 타입은 상호 타입 변환이 가능.

### String 타입을 []byte로 타입 변환할 수 있다.
- string 타입과 []byte 타입은 상호 타입 변환이 가능.
- []byte는 1바이트의 가변 길이배열.
- 메모리는 1바이트 단위로 저장되기 때문에 모든 문자열은 1바이트 배열로 변환이 가능.

## 문자열 순회
문자열을 순회하는 방법은 크게 3가지.
### 인덱스를 사용한 바이트 단위 순회
- 영어 알파벳과 한글이 섞인 문자열 인덱스를 사용해 순회하는 예제
```
func example() {
  str := "Hello 월드!"

  for i :=0; i < len(str); i++ {
    fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
  }
}

 타입:uint8 값:72 문자값:H
 타입:uint8 값:101 문자값:e
 타입:uint8 값:108 문자값:l
 타입:uint8 값:108 문자값:l
 타입:uint8 값:111 문자값:o
 타입:uint8 값:32 문자값: 
 타입:uint8 값:236 문자값:ì
 타입:uint8 값:155 문자값:
 타입:uint8 값:148 문자값:
 타입:uint8 값:235 문자값:ë
 타입:uint8 값:147 문자값:
 타입:uint8 값:156 문자값:
 타입:uint8 값:33 문자값:!

// len() 함수는 바이트를 반환하였고 영어는 1바이트, 한글은 3바이트 이기 때문에 한글이 깨져서 표기됨.
```

### []rune으로 타입 변환 후 한 글자씩 순회
- 한글이 섞인 문자열을 []rune타입으로 변환한 다음 순회
```
func example() {
  str := "Hello 월드!"
  arr := []rune(str)

  for i := 0; i < len(arr); i++ {
    fmt.Printf(" 타입:%T 값:%d 문자값:%c\n", arr[i], arr[i], arr[i])
  }
}

 타입:int32 값:72 문자값:H
 타입:int32 값:101 문자값:e
 타입:int32 값:108 문자값:l
 타입:int32 값:108 문자값:l
 타입:int32 값:111 문자값:o
 타입:int32 값:32 문자값: 
 타입:int32 값:50900 문자값:월
 타입:int32 값:46300 문자값:드
 타입:int32 값:33 문자값:!
 
 // 한 글자씩 순회할 수 있지만 별도의 배열을 할당하므로 불필요한 메모리 사용함.
```

### range 키워드를 이용해 한글자씩 순회
- 한글이 섞인 문자열을 range 키워드를 이용해 순회
```
func example() {
  str := "Hello 월드!"

  for _, v := range str {
    fmt.Printf(" 타입:%T 값:%d 문자:%c\n", v, v, v)
  }
}

타입:int32 값:72 문자:H
 타입:int32 값:101 문자:e
 타입:int32 값:108 문자:l
 타입:int32 값:108 문자:l
 타입:int32 값:111 문자:o
 타입:int32 값:32 문자: 
 타입:int32 값:50900 문자:월
 타입:int32 값:46300 문자:드
 타입:int32 값:33 문자:!

// 정상적으로 순회 확인
```

## 문자열 합치기
문자열은 +과 += 연산을 사용해 이을 수 있다.
- 두 문자열을 잇는 예제
```
func example() {
  str1 := "Hello"
  str2 := "World"

  str3 := str1 + " " + str2
  fmt.Println(str3)

  str1 += " " + str2
  fmt.Println(str1)
}

Hello World
Hello World
```

### 문자열 비교하기
연산자 --, !=를 사용해서 문자열이 같은지 같지 않은지 비교함.
- 예제
```
	str1 := "Hello"
	str2 := "Hell"
	str3 := "Hello"

	fmt.Printf("%s == %s  : %v\n", str1, str2, str1 == str2)
	fmt.Printf("%s != %s : %v\n", str1, str2, str1 != str2)
	fmt.Printf("%s == %s : %v\n", str1, str3, str1 == str3)
	fmt.Printf("%s != %s : %v\n", str1, str3, str1 != str3)

Hello == Hell  : false
Hello != Hell : true
Hello == Hello : true
Hello != Hello : false
```

### 문자열 대소 비교하기
- 연산자 '>', '<' '>=' '<=' 를 이용하여 문자열 간 대소를 비교
- 첫 글자부터 하나씩 값을 비교해서 그 글자에 해당하는 유니코드 값이 다를 경우 대소 반환.
- 예제
```

func example() {
	str1 := "BBB"
	str2 := "aaaaAAA"
	str3 := "BBAD"
	str4 := "ZZZ"

	fmt.Printf("%s > %s : %v\n", str1, str2, str1 > str2)   
	fmt.Printf("%s < %s : %v\n", str1, str3, str1 < str3)   
	fmt.Printf("%s <= %s : %v\n", str1, str4, str1 <= str4) 
}

BBB > aaaaAAA : false // B: 66, a:97
BBB < BBAD : false // B: 66, A: 65
BBB <= ZZZ : true // B:66, Z: 90
```

## 문자열 구조
### String 구조
- Go 언어에서 제공하는 내장 타입으로 내부 구현은 감춰져 있음.
- reflect 패키지 안의 StringHeader 구조체를 통해서 볼 수 있음.
```
type StringHeader struct {
    Data uintptr // 문자열의 데이터가 있는 메모리 주소
    Len int // 문자열의 길이
}
```
### String끼리 대입
- string 변수가 대입할 경우 Data의 포인터값과 Len값이 복사된다.
```
func example() {
	str1 := "Hello World!"
	str2 := str1 // str1 변수값을 str2에 복사

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)) // Data값 추출
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)) // Data값 추출

	fmt.Println(stringHeader1) // 각 필드 값을 출력합니다.
	fmt.Println(stringHeader2)
}

&{4310136699 12}
&{4310136699 12}

// 강제 타입 변환을 위해 unsafe 사용
```

## 문자열은 불변(Immutable)
- string 타입이 가리키는 문자열의 일부만 변경할 수 없다.
```
var str string = "Hello World"
str = "How are you?" // 성공
str[2] = 'a'  // 에러
```
- slice 타입으로 변환할 경우 새로운 메모리 공간에 할당하여 값을 변경됨.
```
func example() {
	var str string = "Hello World"
	var slice []byte = []byte(str) // 슬라이스로 타입 변환

	slice[2] = 'a' // 3번째 문자 변경

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}

Hello World
Healo World
```

### 문자열 합산
- Go 에서는 String 타입 간 합 연산을 지원한다.
- 예제
```
func example() {
  var str string = "Hello"
	stringheader := (*reflect.StringHeader)(unsafe.Pointer(&str)) // 내부 구조체로 변경
	addr1 := stringheader.Data                                    // Data 필드값 addr1 변수로 저장

	str += " World"            // World 문자열 합산
	addr2 := stringheader.Data // Data 필드값 addr2 변수로 저장

	str += " Welcome!"         // Welcome 합산
	addr3 := stringheader.Data // Data 필드값 add3 변수로 저장

	fmt.Println(str)
	fmt.Printf("addr1:\t%x\n", addr1)
	fmt.Printf("addr2:\t%x\n", addr2)
	fmt.Printf("addr3:\t%x\n", addr3)
}

addr1:  10480bf47
addr2:  14000126002
addr3:  1400012c000

// 모두 다른 메미로 주솟값 -> 기존 메모리 공간은 그대로 유지.
```
- strings 패키지의 Builder를 사용하여 메모리 낭비를 줄일 수 있음.
```
func example() {
  var builder strings.Builder

  // 문자열을 추가
  builder.WriteString("Hello, ")
  builder.WriteString("World!")

  // 문자열 빌드 결과 가져오기
  result := builder.String()

  // 빌드된 문자열 출력
  fmt.Println(result)
}
Hello, World!

```
### 문자열 불변 원칙을 지키는 이유
- 예기치 못한 버그를 방지
- String 타입이 복살될 때 문자열 전체가 아닌 Data, Len 필드값만 복사되기 때문에.
- 스레드 안정성, 성능 최적화, 참조 안정성 등..
