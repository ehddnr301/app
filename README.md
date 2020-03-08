# Golang

- go run main.go

- how to export function = start uppercase

### Variable

- ````const name string = "hello"

  var studying bool = true
  studying = false

  shortcut := "good"
  shortcut = "excellent"

  fmt.Println(name, studying, shortcut) ```
  ````

- 변수를 만들때 type을 적어줘야하지만 shortcut을 사용하면 알아서 지정한다.
- function 안에서는 shortcut 처럼 쓸 수 있다.

### Function

- ```func lenAndUpper(name string) (int, string) {
  )
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
