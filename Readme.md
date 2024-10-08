# Go Language
## Introduction
- Go is an open-source programming language developed by Google. It is a statically typed, compiled language with a syntax similar to C. Go is designed for simplicity, efficiency, and performance. It is widely used for building web applications, cloud services, and system software.

## Basics
   ### Package
    - A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
In go it must have a main package ie :
```go
package main
```
   ### Import
    - The import keyword is used to include the code from other packages to use in our program. 
```go
import "fmt"
```
- Note : Fmt is a standard library package that provides formatting and printing functions.
   ### Function
    - A function is a block of code that performs a specific task. It has a name, a list of parameters, an optional return type, and a body. 
```go
func main() {
    fmt.Println("Hello, World!")
}
```
   ### Variable
    - A variable is a storage location in memory that is represented by a name, a type, and a value. 
```go
var name string = "John"
```
   ### Constant
    - A constant is a variable whose value cannot be changed once it is assigned. 
```go
const pi = 3.14
```
   ### Data Types
    - Go is a statically typed language, which means that the type of a variable is known at compile time. 
```go
var name string = "John"
var age int = 30
var isMarried bool = true
var height float64 = 5.8
```

- Note : There is no necessity of defining the data type explicitly, Go can infer the data type from the value assigned to the variable. But in some cases the data type must be defined explicitly such as when the variable is not initialized with a value.
 - Note: In go , variables can be defined in more shorter way as :
```go
name := "John"
age := 30
isMarried := true
height := 5.8
```
   ### Operators
    - Go supports a variety of operators for performing arithmetic, logical, and comparison operations. 
```go
var sum = 10 + 20
var diff = 30 - 10
var product = 5 * 6
var quotient = 10 / 2
var remainder = 10 % 3
```
# For Loop 
- The for loop is used to execute a block of code repeatedly until a specified condition is met. 
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

Note : There is no while loop in go , but we can use for loop as while loop as :
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```
And also we can use for loop as infinite loop as :
```go
for {
    fmt.Println("Hello, World!")
}
```
And there is concpt of range in for loop as :
```go

for i := range 3 {
		fmt.Println(i)
	}
```
# If Else
- The if statement is used to execute a block of code if a specified condition is true. 
```go
if age >= 18 {
    fmt.Println("You are an adult")
} else {
    fmt.Println("You are a minor")
}
```
Note: The else statement is optional, and the else if statement can be used to check multiple conditions.
```go
if age >= 18 {
    fmt.Println("You are an adult")
} else if age >= 13 {
    fmt.Println("You are a teenager")
} else {
    fmt.Println("You are a child")
}
```
And we can also define the variable in if statement as :
```go
if age := 30; age >= 18 {
    fmt.Println("You are an adult")
}
```

# Switch Case
- The switch statement is used to execute a block of code based on the value of a variable. 
```go
switch day {
    case "Monday":
        fmt.Println("Today is Monday")
    case "Tuesday":
        fmt.Println("Today is Tuesday")
    default:
        fmt.Println("Today is a different day")
}
```

Note : The switch statement can also be used without a condition to create a switch case statement.
```go
switch {
    case day == "Monday":
        fmt.Println("Today is Monday")
    case day == "Tuesday":
        fmt.Println("Today is Tuesday")
    default:
        fmt.Println("Today is a different day")
}
```

Also it can be used for type switch as :
```go
var x interface{} = 10

switch x.(type) {
    case int:
        fmt.Println("x is an integer")
    case string:
        fmt.Println("x is a string")
    default:
        fmt.Println("x is another type")
}
```
# Arrays
- An array is a collection of elements of the same type stored in contiguous memory locations.
```go
var numbers = [5]int{1, 2, 3, 4, 5}
```
- Note : The size of the array is fixed and cannot be changed once it is defined.
- Note : The elements of the array can be accessed using the index of the element.
```go
numbers[0] = 10
```
- Note : The length of the array can be obtained using the len() function.
```go
len(numbers)
```
 Array can be declared in single line as :
```go
numbers:=[5]int{1, 2, 3, 4, 5}
```
And Two dimnesional array can be declared as :
```go
numbers:=[2][3]int{{1, 2, 3}, {4, 5, 6}}
```



# Slices
- A slice is a dynamically-sized, flexible view into a portion of an underlying array.
```go
var numbers = []int{}
```
- Note : The length of the slice can be obtained using the len() function.
```go
len(numbers)
```
- Note : The capacity of the slice can be obtained using the cap() function.
```go
cap(numbers)
```
- Note : The elements of the slice can be accessed using the index of the element.
```go
numbers[0] = 10
```
- Note : The append() function is used to add elements to the slice.
```go
numbers = append(numbers, 5)
```
- Note: The make() function is used to create a slice with a specific capacity.
```go
numbers = make([]int, 5)
```
It takes three arguments: the first argument is the type of the slice, the second argument is the length of the slice, and the third argument is the capacity of the slice.

- Note : The copy() function is used to copy elements from one slice to another.
```go
numbers2 := make([]int, len(numbers))  
copy(numbers2, numbers) //destination, source   
```
Slice Operator:
```go
numbers := []int{1, 2, 3, 4, 5}
numbers2 := numbers[0:3] //  return [1, 2, 3] excluding the last element ie 3rd index
numbers3 := numbers[2:] //  return [3, 4, 5]
numbers4 := numbers[:3] //  return [1, 2, 3]
```
Slice Package :
```go
import "golang.org/x/exp/slices"
var num1 = []int{1, 2, 3}
var num2 = []int{4, 5, 6}
slices.Join(num1, num2)// returns [1, 2, 3, 4, 5, 6]
slices.Equal(num1, num2)// returns  true or false
slices.Clone(num1)// returns [1, 2, 3]
```