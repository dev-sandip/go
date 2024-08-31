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