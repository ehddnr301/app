# Golang

- go run main.go

- how to export function = start uppercase

### Variable

- ```
  const name string = "hello"

  var studying bool = true
  studying = false

  shortcut := "good"
  shortcut = "excellent"

  fmt.Println(name, studying, shortcut)
  ```

- 변수를 만들때 type을 적어줘야하지만 shortcut을 사용하면 알아서 지정한다.
- function 안에서는 shortcut 처럼 쓸 수 있다.

### Function

- ```
  func lenAndUpper(name string) (int, string) {
    return len(name), strings.ToUpper(name)
  }

  func consoleLogging(name ...string) {
  	fmt.Println(name)
  }

  func main() {
  	totalLen, uppername := lenAndUpper("Someting")
  	totalLength, _ := lenAndUpper("sssssssss")
  	fmt.Println(totalLen, uppername, totalLength)

  	consoleLogging("super", "suppppper", "many", "arguments")
  }
  ```

- function은 argument type과 return type을 적어줘야 한다.
- go 는 return 값이 하나가 아닐수도 있다.
- 여러개를 리턴해줄때 무시하는 방법은 underScore 이다.
- 제한없이 arguments를 받으려면 type앞에 ... 을 붙여준다.

### Function - Naked return, defer

- ```
  func lenAndUpper(name string) (length int, upper string) {
    defer fmt.Println("It's Done")
  	length = len(name)
  	upper = strings.ToUpper(name)
  	return
  }

  func main() {
  	totalLen, uppername := lenAndUpper("Someting")
  	fmt.Println(totalLen, uppername)
  }
  ```

- return 할 타입에 미리 변수명을 적어주면 선언이 된다.
- 그리고 함수안에서 변수를 업데이트 한다.
- return 을 적어주면 알아서 해당 변수를 리턴해준다.
- defer는 함수가 끝난뒤에 실행된다.

### For, range

- ```
  func superAdd(numbers ...int) int {
    	total := 0
    	for _, number := range numbers {
    		total += number
    	}
    	return total
  }

  func main() {
  	fmt.Println(superAdd(1, 2, 3, 4, 5, 6, 7))
  }
  ```

- for 에서 첫번째 인자는 Index라서 underScore로 ignore시킴.

### if문

- ```
  func canIDrink(age int) bool {
    	if kAge := age + 2; kAge < 20 {
    		return false
    	}
    	return true
  }

  func main() {
  	fmt.Println(canIDrink(20))
  	fmt.Println(canIDrink(17))
  }
  ```

- if문 안에서 변수를 생성할수있다.

### Switch

- ```
    func canIDrink(age int) bool {
      	switch koreanAge := age + 2; koreanAge {
      	case 10:
      		return false
      	case 19:
      		return true
      	}
      	return false
    }

  func main() {
  	fmt.Println(canIDrink(20))
  	fmt.Println(canIDrink(17))
  }
  ```

- if문과 비슷하게 변수생성 가능.

### Pointer

- ```
      func main() {
          a := 2
        	b := a
        	c := &a
        	a = 3
        	fmt.Println(a, b, *c)
      }
  ```
- a는 2
- b는 (a의 값) 2
- a는 3
- 따라서 a는 3 b는 2가 출력된다.
- `fmt.Println(&a, &b)` 을 하면 메모리 주소를 출력할수있다.
- `*` 는 see through 라는 의미 훑어본다는 의미.
- 따라서 c는 a의 메모리 주소를 나타내지만 \*로 훑어 봤기때문에 a가 3으로 변경되었으니까 \*c는 3이 찍힌다.

- ```
    func main() {
    	a := 2
    	c := &a
    	*c = 3
    	fmt.Println(a)
    }
  ```
- a는 2
- c는 a의 메모리 주소 와 연결되어있음.
- c의 훑어본 값은 3
- a는 3

- 요약 : `*`는 메모리 주소에 저장된 값 `&`는 메모리 주소

### array, slice

- ```
    func main() {
    	names := [5]string{"as", "ax", "ew"}
    	supers := []string{"as123", "ax123", "ew123"}
    	supers = append(supers, "superman")
    	fmt.Println(names, supers)
    }
  ```

- array는 기본적으로 길이를 제한해 주어야 한다.
- 제한없는 array를 slice라고 부른다.
- slice에 추가하려면 append를 사용해야한다.
- append는 첫번째로 slice, 두번째로 추가할것 을 받는다.
- append는 새로운 slice를 리턴한다.

### map

- ```
  func main() {
  	superman := map[string]string{"name": "super", "age": "12"}
  	for key, value := range superman {
  		fmt.Println(key, value)
  	}
  }
  ```
- map 은 object같은것인데 key의 데이터타입과 value의 데이터 타입을 적어줘야한다.
- for range도 사용가능하다.

### structs

- ```
    type person struct {
    	name    string
    	age     int
    	favFood []string
    }

    func main() {
    	favFood := []string{"kimchi", "삼겹살"}
    	superman := person{"suuper", 23, favFood}
    	batman := person{name: "baat", age: 50, favFood: favFood}
    	fmt.Println(superman, batman)
    }
  ```

- struct 가 object와 class 같은 것인데 value타입을 다르게 적어줄수있다.
- batman과 같은 방법을 쓰는것이 가독성이 좋다.
