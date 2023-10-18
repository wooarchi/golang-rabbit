# Pointer

## 포인터란?
포인터는 메모리 주소를 가리키는 변수이며 변수의 주소를 직접 참조하고 조작할 수 있다.
- 포인터 변수 정의 및 값 출력
```
  var a int = 3 // int 변수 선언
  var p *int // pointer 변수 선언

  p = &a

  fmt.Println(p) // 0x140000a2008
```

- 여러 포인터 변수가 하나의 메모리 공간을 가리킬 수 있다.
```
  // int 변수 선언
  var a int = 3 

  // 포인터 변수 p1, p2, p3 선언
  var p1 *int = &a
  var p2 *int = &a
  var p3 *int = &a

  fmt.Println(p1) // 0x140000a2008
  fmt.Println(p2) // 0x140000a2008
  fmt.Println(p3) // 0x140000a2008
```

- 같은 메모리 공간을 가리키는 포인터 변수의 값은 동일하게 변화한다.
```
  // int 변수 선언
  var a int = 3 

  // 포인터 변수 p1, p2, p3 선언
  var p1 *int = &a
  var p2 *int = &a
  var p3 *int = &a


  fmt.Println(*p1) // 3
  fmt.Println(*p2) // 3
  fmt.Println(*p3) // 3
  
  var b int = 4

  *p1= b

  fmt.Println(*p1) // 4
  fmt.Println(*p2) // 4
  fmt.Println(*p3) // 4
```

## 포인터 변수 선언
- 포인터 변수는 가리키는 데이터 타입 앞에 *을 붙여서 선언
```
# int 타입 데이터의 메모리 주소를 가리키는 포인터 변수
var p *int

# float64
var p *float64

# User 구조체
var p *User
```

- 포인터 변수 사용 예제
```
func example() {
  var a int = 500
  var p *int

  p = &a

  fmt.Printf("p의 값: %p\n", p)
  fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

  *p = 100
  fmt.Printf("a의 값: %d\n", a)

}
p의 값: 0x140000a4008
p가 가리키는 메모리의 값: 500
a의 값: 100
```

- 포인터 변숫값 비교
```
// == 연산을 사용해 같은 메모리 공간을 가리키는지 확인

func example() {
  var a int = 10
  var b int = 20

  var p1 *int = &a
  var p2 *int = &a
  var p3 *int = &b

  fmt.Printf("p1 == p2 : %v\n", p1 == p2)
  fmt.Printf("p2 == p3 : %v\n", p2 == p3)

}
p1 == p1 : true
p2 == p3 : false
```
- 포인터의 기본값운 nil
```
func example() {
  var p *int

  fmt.Println(p)
}

<nil>
```

## 포인터를 사용하는 이유
- 변수 대입, 함수 인수 전달은 항상 값을 복사하기 때문에 많은 메모리 사용.
- 다른 공간으로 복사되기 때문에 변경사항이 적용되지 않음.


### Data 구조체를 생성해 포인터 변수 초기화
- 기존방식
```
var data Data
var p *Data = &data
```

- 구조체를 생성해 초기화하는 방식
```
var p *Data = &Data{}
```

## 인스턴스
인스턴스란 메모리에 할당된 데이터의 실체.
```
// Data 타입값을 저장할 수 있는 메모리 공간 할당: 인스턴스
var data Data

// Data 타입 포인터 변수 선언
var p1 *Data = &data // 현재 인스턴스는 1개
var p2 *Data = p1 // 현재 인스턴스는 1개
var p3 *Data = p2  // 현재 인스턴스는 1개

// Data 타입 변수 선언
var data1 Data // 현재 인스턴스 1개
var data2 Data = data1 // 현재 인스턴스 2개
var data3 Data = data1  // 현재 인스턴스 3개
```

### 인스턴스는 데이터의 실체다.
- 인스턴스는 메모리에 존재하는 데이터의 실체.
- 포인터를 사용하여 인터스턴스에 접근.
- 구조체 포인터를 함수 매개변수로 받음.

### new() 내장 함수
- 포인터값을 별도의 변수를 선언하지 않고 초기화.
```
var p1 := &Data{}
var p2 := &Data{3, 4}
```
- new()를 사용하여 초기화: 타입을 메모리에 할당하고 기본값을 채워 주소 반환.
```
var p2 = new(data)
```

### 인스턴스가 사라지는 시점
- 메모리는 무한한 리소스가 아니기 때문에 불필요한 데이터를 메모리에서 제거하는 기능이 필요하다.
- Go 언어는 'Garbage Collector' 기능을 사용하여 메모리에서 불필요한 데이터를 제거한다.
- 예제
```
func TestFunc() {
	u := &User{} // u 포인터 변수 생성(인스턴스 생성)
    u.Age = 30
    fmt.Println(u)
} // 내부 변수가 사라짐 -> 인스턴스도 사라짐.
```

## 스택 메모리와 힙 메모리
- 스택 메모리 영역이 힙 메모리 영역보다 훨씬 효율적.
- 스택 메모리는 함수 내부에서만 사용 가능.
- 함수 외부로 공개되는 메모리 공간은 힙 메모리 영역에서 할당.
- Go 언어는 'Escape analysis'를 사용하여 어느 메모리에 할당할지 결정
- 예제(함수 외부로 공개되는 인스턴스의 경후 함수가 종료되어도 사라지지 않음)
```
type User struct {
	Name string
    Age  int
}

func NewUser(name string, age int) *User {
	var u = User{name, age}
    return &u // 탈출 분석을 통해 메모리가 사라지지 않는다.
}

func main() {
	userPointer := NewUser("Coen", 34)
    
    fmt.Println(userPointer)
}

&{AAA 23}
```
