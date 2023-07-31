### 2.1.2 literals

* Go에서 찾아볼 수 있는 literal은 일반적으로 네 가지가 있다.
    * 드물게 사용되는 5번째 literal은 complex number를 다룰 때 볼 수 있다.

#### integer literals

* 일반적으로 10진수를 의미. 접두사를 사용하여 진법을 변경하여 사용할 수 있다.
    * `0b`: 2진수, `0o`: 8진수, `0x`: 16진수
    * 접두사로 대문자, 소문자를 사용할 수도 있다
    * 문자 없이 0으로 시작하는 숫자는 8진 literal을 표한하는 또 다른 방법이지만 혼동이 있을 수 있어 사용하지 않는게 좋다.

* 긴 integer literal을 쉽게 읽기위해 Go는 integer literal 사이에 밑줄을 넣을 수 있다.
* 1000의 단위로 구분하기 위해서 `1_234` 같이 사용할 수 있다.
    * 숫자의 맨 앞이나 맨 뒤에 넣을 수 없고, 다른 숫자와 연결되게 사용해서는 안된다.
* 10진수를 천 단위로 나누거나, 2진수, 8진수, 16진수를 1, 2, 4 바이트 경계로 나누어 표시하여 가독성을 높이는 용도로 사용한다.

#### floating point leterals

* 값에서 소수부를 구분하는 소수점이 있다.
* 문자 e와 양수 혹은 음수(6.03e23과 같은)로 지정된 지수를 가질 수 있다.
* 16진수로 floating point를 표현하기 위해 0x Prefix와 지수를 나타내는 p를 사용할 수 있다.
* integer literal과 마찬가지로 formatting하기 위해 밑줄을 사용할 수 있다.

#### rune literals

* 문자를 나타내며, single quotation marks로 묶어 사용한다.
* Go에서 single quotation marks와 double quotation marks는 혼용할 수 없다.

### 2.1.4 Numeric types

#### specific integer type

* `byte` type은 `uint8`의 alias이다. 웬만하면 `uint8` 대신에 `byte`를 사용하자.

> int type을 32bit 부호 있는 정수로 사용하는 일반적이지 않은 64 bit architecture가 있다.
> Go는 이중 3개를 지원한다. `amd64p32`, `mips64p32`, `mips64p32le`

* 정수의 부호가 없는 `uint` type

### 2.2 var vs. :=

* 함수 내에서 `:=` 연산자 사용을 피해야하는 상황
    * zero value으로 변수를 initialize할 때, `var x int`와 같이 사용하여, 의도한 zero value이 할당될 수 있도록 하기 위함이다.
    * type이 지정되지 않은 constant나 literal이 변수에 할당될 때, type을 지정할 수 있는 var 형식을 사용하자. type 변환을 사용하여 값의 타입을 지정하고 `:=` 연산자를 사용하여 `x := byte(20)`과 같이 사용하는 것은 문제는 없지만, `var x byte = 20`으로 사용하는 것이 관용적인 코드를 만들 수 있다.
    * `:=` 연산자는 새로운 변수나 이미 존재하는 변수에 모두 값을 할당하기 때문에, 이미 있는 변수에 값을 할당하려고 의도했던 것이 새로운 변수를 만들게 되는 경우가 발생한다. 이런 상황에서 명시적으로 var keyword를 사용하여 변수를 새로운 선언으로 명확히 한다면, 이미 있던 변수나 새로운 변수 모두에 할당 연산자(=)를 사용할 수 있다.


> 데이터 흐름 분석을 하기 어려워질 수 있는 함수 외부의 변수 선언은 피하도록 하자.

### 2.3 Using const

> Go에서 constant는 literal에 이름을 부여하는 방법이다. Go에서 변수를 변경 불가능하게 선언하는 방법은 없다.

### 2.4 Typed and Untyped Constants

* Constant는 type이 지정되거나 지정되지 않게 할 수 있다. type이 지정되지 않은 const는 literal과 똑같이 처리된다. 자체 type은 없지만, 다른 type으로 추론할 수 없을 때 사용하는 기본 type을 갖고 있다. typed const는 해당 type의 변수에 직접적으로 할당될 수 있다.

### 2.5 Unused Variables

```go
func main(){
    x := 10
    x = 20
    fmt.Println(x)
    x = 30
}
```

* compiler와 `go vet`은 사용되지 않은 10과 30의 할당을 잡아내지 못하지만 `golangci-lint`는 잡아낼 수 있다.

> Go compiler는 접근하지 않는 package level 변수 생성을 내버려둘 것이다. 이는 package level 변수 만드는 것을 피해야 하는 이유다.

### 2.6 Naming Variables and Contants

* 관용적인 Go는 `snake case`를 사용하지 않기 때문에 밑줄 사용은 드물다. 대신에 식별자 이름이 여러 단어로 구성된 경우에 `camel case`를 사용한다.

> Go에서 밑줄(_) 자체는 특별한 식별자 이름이다.

* Go는 package level 선언의 이름에서 첫번째 문자의 대소문자 사용으로 package 외부에서 접근 가능한지 여부를 결정하기 때문에, 상수를 항상 대문자로 쓴다거나 하는 방식을 사용하지 않는다.
* 함수 내에서는 짧은 변수 이름을 선호한다. 예를 들어, `k`와 `v`(각각 `key`와 `value`의 줄임)는 `for-range loop` 에서 변수 이름으로 사용된다.
    * 짧은 이름은 두가지 용도로 사용된다. 첫번째는 반복적인 입력을 제거하여 코드를 간결하게 하는 것이고, 두번째는 코드가 얼마나 복잡한 상태인지 확인하는 역할을 한다. 짧은 이름의 변수를 추적하기 힘들다고 판단된다면, 해당 code block이 너무 많은 일을 한다는 의미이다.
* package block 내에 constants 와 variables 이름을 만들 때, 더 설명이 가능한 이름으로 사용하자. type은 변수 이름에서 여전히 제외되어야 하지만 범위가 더 넓어지면, 명확하게 하는 더 완벽한 이름이 필요하다.