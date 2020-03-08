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
