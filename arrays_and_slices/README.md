# Arrays
```go
	// create an array called "firstArray" that has length 3
	// and holds integers
	firstArray := [3]int{1, 2, 3}
	fmt.Printf("firstArray: %v, type: %T", firstArray, firstArray)

  // If initializing on declaration, this is allowed
  // Creates an array just large enough to hold data
  firstArray := [...]int{1, 2, 3}

  var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Dylan"
	students[2] = "Tom"
	fmt.Printf("Students: %v", students)
	fmt.Printf("Students[1]: %v", students[1])
  fmt.Printf("len(students): %v\n", len(students))

  // Lets make the identity matrix!
  var identityMatrix [3][3]int = [3][3]int{[3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1}}

  // or 
  var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

  // or
  var identityMatrix [3][3]int
  identityMatrix[0] = [3]int{1,0,0}
  identityMatrix[1] = [3]int{0,1,0}
  identityMatrix[2] = [3]int{0,0,1}
```

In **go** arrays are actually considered **values**. In other languages arrays are **pointing** to values in memory. This is incredibly important when copying an array in **go**, as you are actually creating an entirely *new* value in memory. So if you change the copied array the original array will not be affected.
```go
a := [...]int{1,2,3}
b := a
b[1] = 5
fmt.Println(a) //[1 2 3]
fmt.Println(b) //[1 5 3]
```
What if we want to point to an existing array instead of making a new copy, in this case we use the `&` operator!
```go
a := [...]int{1,2,3}
b := &a
b[1] = 5
fmt.Println(a) //[1 5 3]
fmt.Println(b) //&[1 5 3]
```

# Slices

Whats a slice look like?
```go
a := []int{1,2,3}
```
Looks the same as an array but instead of `[...]int{1,2,3}` we have `[]int{1,2,3}` ! We can actually do everything to a **slice** that we can do to an **array**.
```go
a := []int{1,2,3}
fmt.Println(len(a)) //3
fmt.Println(cap(a)) //3
```
Futhermore, slices naturally refer to the underlying data.  If multiplce slices refer to the same underlying data, changing one slice changes them all.

```go
a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
b := a[:]   // slice from all elements
c := a[3:]  // slice from 4th element to end
d := a[:6]  // slice first 6 elements
e := a[3:6] // slice 4th, 5th, 6th
// NOTE: FIRST NUMBER IS INCLUSIVE, SECOND IS EXCLUSIVE!!!
fmt.Println(a) 
fmt.Println(b) 
fmt.Println(c) 
fmt.Println(d) 
fmt.Println(e) 
// [1 2 3 4 5 6 7 8 9 10]
// [1 2 3 4 5 6 7 8 9 10]
// [4 5 6 7 8 9 10]
// [1 2 3 4 5 6]
// [4 5 6]
```

The last way to make slices is using the built in `make` function.
```go
a := make([]int, 3, 100)
fmt.Println(a)      //[0,0,0]
fmt.Println(len(a)) //3
fmt.Println(cap(a)) //100
```

Adding elements to slice
```go
a := []int{}
fmt.Println(a)      //[]
fmt.Println(len(a)) //0
fmt.Println(cap(a)) //0
a = append(a, 1)
fmt.Println(a)      //[1]
fmt.Println(len(a)) //1
fmt.Println(cap(a)) //1
a = append(a, 2, 3, 4, 5)
fmt.Println(a)      //[1 2 3 4 5]
fmt.Println(len(a)) //5
fmt.Println(cap(a)) //6
// concat slices
a = append(a, []int{2, 3, 4, 5}...)
```

remove last element from slice
```go
  a := []int{1, 2, 3, 4, 5}
  b := [:len(a)-1]
  fmt.Println(b)
```
remove first element from slice
```go
	a := []int{1, 2, 3, 4, 5}
	b := a[1:]
	fmt.Println(b)
```