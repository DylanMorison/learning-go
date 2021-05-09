5/8/2021
# Learning Go - Chapter 1

**Basic Installation**
```
# use this to reat a go program like a script
# and run source code.
go run hello.go

# Create a binary that is distributed for other
# people to use
go build hello.go

# compile to a binary called "hello_world"
go build -o hello_world.go
```

**goimports** install
```
# cleans up go import statements
go install golang.org/x/tools/cmd/goimports@latest
```

Variable declaration
```golang
//declare
var foo int

//declare and initialize
var foo int = 42

//declare and initialze but let compiler figure out type
//(shorthand for above)
foo := 42
```

Can't redeclare variable, but can shadow them
```go
var (
	helloString string = "hi"
	helloInt    int    = 1
)

func main() {
  //need to shadow variables from within main()
  var (
	  helloString string = "hi"
	  helloInt    int    = 1
  )
}
```

All variables must be used!

Visibility
- lower case first letter for package scope
- upper case first letter to export
- no private scope

Naming conventions:
- Pascal or camelCase
- Capitalize acronyms (HTTP, URL)
- As short as reasonable
  - longer name for longer lives

Type conversion example:
```go
package main
// ! Go does NOT to implicit type conversions!
// use strconv package to convert to str
import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
  // convert int to str
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)

  // int to float32
	var testInt int = 1
	var intToFloat float32
	fmt.Printf("%v, %T\n", testInt, testInt)
	intToFloat = float32(testInt)
	intToFloat = 1.5
	fmt.Printf("%v, %T\n", intToFloat, intToFloat)
}
```
