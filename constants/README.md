
## Summary
- Immutable, but can be shadowed
- Replaced by the compiler at compile time
  - Value must be calculable at compile time

Named Like Variables
- PascalCase for exported constants
- camelCase for internal constants
```go
package main

import "fmt"

func main() {
  // if we are going to export change myConst to MyConst
	const myConst int = 42
}
```

Typed Constants work like immutable variables
- Can interoperate only with same type

Untyped Constants work like immutable variables
- Can interoperate only with same type

Enumerated Constants
- Special symbol iota allows related constants to be created easily
- Iota starts at 0 in each const block and increments by one
- Watch out for constant values that match zero values for variables

Enumerated Expressions
- Operations that can be determined at compile time are allowed, such as
  - Arithmetic
  - Bitwise 
  - Bitshifting