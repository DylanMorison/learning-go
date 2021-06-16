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

Variable Types
- Boolean
  - T or F
  - Not an alias for other types (e.g. int)
  - Zero or False value
- Numeric Types
  - Integers
    - Signed integers
      - int type has varying size, but min 32 bits
      - 8 bit (int 8) through 64 bit (int 64)
    - Unsigned integers
      - 8 bit (byte and unit8) through 32 bit (uint32)
    - Arithmetic Operations
      - Addition, subtraction, multiplication, division, remainder
    - Bitwise operations
      - And, or, xor, and not
    - Cant mix types in same family! (uint16+uint32 = error)
  - Floating Point Numbers
    - Following IEEE-754 standard
    - 32 and 64 bit versions
    - Literal Styles
      - Decimal (3.14)
      - Exponential (13e18)
      - Mixed (13.7e12)
    - Arithmetic Operations
      - Addition, subtraction, mult, div
  - Complex Numbers
      - Zero value is `0+0i`
      - 64 and 128 bit versions
      - Built in functions
        - complex: makes complex num from two floats
        - real: get real part as float
        - imag: get imaginary part as float
      - Arithmetic Operations
      - Addition, subtraction, mult, div
- Text Types
  - Strings
    - UTF-8
    - Immutable
    - Can be concatenated with plus (+) operator
    - Can be converted to []byte
  - Rune
    - UTF-32
    - Alias for int32
    - Special methods normally required to process
      - e.g. strings.Reader#ReadRune




