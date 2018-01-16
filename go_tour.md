# Go Tour Notes
## Commands
`echo $GOPATH` - Check if the path is pointing to the workspace.
`cd $GOPATH` - Where you store all your Go files.
`go fmt` - A command to format code file.
`fmt` - A library/Package used print a line of code.
`go run _filename_.go` - Only runs the Go code in a specific file.
`go install` - Complies all Go source code in a folder into machine code and saves it in the /bin folder. You can run the code (even if you delete the original GO folder and/or file) by typing the folder name in the terminal.

## Basics
### Packages, Imports
``` go
// Go program is made up of packages and start running in `main`
// the package that is responsible for execution of a program
package main

// last element of import path is the package name
import (
	"fmt"
  "math/rand"
)

// he main package cannot be built without the existence of a main func. Unlike dynamic typed, scripting languages (e.g. Ruby, Python, JavaScript and etc...), Go does not allow you to execute any function outside of main.
// exported names begins with capital letter e.g. `.Println`, `.Intn`
func main() {
	rand.Seed(10)
	fmt.Println("My favorite number is", rand.Intn(10))
}
// => My favorite number is 4

func main() {
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
// => Now you have 2.6457513110645907 problems.
```

### Functions

``` go
package main

import "fmt"

// An argument has a variable and type. If arugments have the same type, only add type to the last parameter.
// return value needs a type
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
// => 55

//_______________Multiple results______________________
// more than one type for the return value
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
// => "world hello"

//_______________Naming Return Values__________________
// A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
// => 7 10
```

### Variables
Variables declared without an explicit initial value are given their zero value.

The zero value is:
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

``` go
// Declares a list of variables, add type at the end.
var i, j int

// Variables with initializers
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var c, python, java = true, false, "no!"

// Short variable declarations can only be used inside a Function
func main() {
	k := 3
	fmt.Println(k)
}
// => 3
```

### Types
``` go
// _________Basic types in Go____________
bool
string // every char in a string is a byte (8bit), a string is an array of uint8

// (2^x - 1) where x = int_ gives you have much space you can use
int  int8  int16  int32  int64 // signed int e.g. int4 => _000 (first bit is used for sign +/-) => #'s from -7 to 7
uint uint8 uint16 uint32 uint64 uintptr // unsigned int e.g uint4 => 0000 => #'s from 0 to 15

byte // alias for uint8
rune // alias for int32
     // represents a Unicode code point
float32 float64 // int with decimals
complex64 complex128

// ___________Type conversions___________
// The expression T(v) converts the value v to the type T.
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
Or, put more simply:
i := 42
f := float64(i)
u := uint(f)

// ____________Type inference___________
// When the right hand side of the declaration is typed, the new variable is of that same type:
var i int
j := i // j is an int
// But when the right hand side contains an untyped numeric constant, the new variable may be an int, float64, or
// complex128 depending on the precision of the constant:
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

### Constants
Constants are declared like variables, but with the const keyword.
Constants cannot be declared using the := syntax.
``` go
package main

import "fmt"

func main() {
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
// => Go rules? true
```

### Control statements
``` go
package main

import (
	"fmt"
	"math"
	"time"
)

//___________for______________
// For is the only looping construct
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// The init and post statement are optional
func main() {
	sum := 1
	for sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}


// _____________if/else__________
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
// => 27 >= 20
// => 9 20

// ________swtich__________
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
  // Switch without a condition is the same as switch true.
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// ____________Defer_____________
// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// => counting
// => done
// => 9
// => 8
// => 7
// => 6
// => 5
// => 4
// => 3
// => 2
// => 1
// => 0
```
### Pointers
A pointer holds the memory address of a value.
``` go
package main

import "fmt"

func main() {
	i, j := 42, 2701
  // The & operator generates a pointer to its operand.
	p := &i         // point to i
	// The type *T is a pointer to a T value. Its zero value is nil.
	fmt.Println(*p) // read i through the pointer
  // The * operator denotes the pointer's underlying value. This is known as "dereferencing" or "indirecting".
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
// => 42
// => 21
// => 73
```

### Structs
https://gobyexample.com/structs
A struct is a collection of fields.
A composite data type declaration that defines a physically grouped list of variables to be placed under one name in a block of memory, allowing the different variables to be accessed via a single pointer, or the struct declared name which returns the same address.
``` go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
  // Struct fields are accessed using a dot.
	v.X = 4
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)

  // Struct fields can be accessed through a struct pointer.
  // To access the field X of a struct when we have the struct pointer p we can just write p.X instead of (*p).X.
	p := &v
	p.X = 10
	fmt.Println(v)
}
// => {1 2}
// => {4 2}
// => 4
// => 2
// => {10 2}


var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

// A struct literal denotes a newly allocated struct value by listing the values of its fields.
func main() {
	fmt.Println(v1, p, v2, v3)
}
// => {1 2} &{1 2} {1 0} {0 0}
```

### Arrays
``` go
// The type [n]T is an array of n values of type T.
package main

import "fmt"

func main() {
	// This expression declares a variable a as an array of two strings
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
// => Hello World
// => [Hello World]
// => [2 3 5 7 11 13]
```

### Slices
A slice is a dynamically-sized, flexible view into the elements of an array.
The type []T is a slice with elements of type T.
Ranges includes the first element, but excludes the last one.
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
``` go
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:4]
	fmt.Println(s)
}
//  => [2 3 5 7]

// When slicing, you may omit the high or low bounds to use their defaults instead.
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)
  // s reference the previous s
	s = s[:2]
	fmt.Println(s)
	// s reference the previous s
	s = s[1:]
	fmt.Println(s)
}
// => [3 5 7]
// => [3 5]
// => [5]

// _________Slice length and capacity__________
// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// => len=6 cap=6 [2 3 5 7 11 13]
// => len=0 cap=6 []
// => len=4 cap=6 [2 3 5 7]
// => len=2 cap=4 [5 7]

// ______Nil slices_____
// The zero value of a slice is nil.
// A nil slice has a length and capacity of 0 and has no underlying array.

// ______Creating a slice with make_______
// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
// The make function allocates a zeroed array and returns a slice that refers to that array
// To specify a capacity, pass a third argument to make
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

// _______Slices of slices___________
board := [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}

// _________Appending to a slice_________
// If the backing array of s is too small to fit all the given values a bigger array will be allocated.
// The returned slice will point to the newly allocated array.

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
// => len=0 cap=0 []
// => len=1 cap=2 [0]
// => len=2 cap=2 [0 1]
// => len=5 cap=8 [0 1 2 3 4]

// __________Range_____________
// When ranging over a slice, two values are returned for each iteration.
// The first is the index, and the second is a copy of the element at that index.

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  // You can skip the index or value by assigning to _.
	for i, v := range pow {
		fmt.Printf("%d %d\n", i, v)
	}
}
// => 0 1
// => 1 2
// => 2 4
// => 3 8
// => 4 16
// => 5 32
// => 6 64
// => 7 128
```
### Maps
The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use.
``` go
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
//  => map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
```
### Methods
Go doesn't have classes.
You can define methods on types.
Methods is a fn with a receiver argument.
https://stackoverflow.com/questions/8263546/whats-the-difference-of-functions-and-methods-in-go
```go
type Vertex struct {
	X, Y float64
}

// The receiver appears in its own argument list between the func keyword and the method name.
//  Abs method has a receiver of type Vertex named v.
// when declaring a method with a receiver, it has to be the same type.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

// method as a fn
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

// Type does not have to be a struct! e.g. `type MyFloat float64`
// when declaring a method with a receiver, it has to be the same type.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	a := MyFloat(-math.Sqrt2)
	fmt.Println(a.Abs())
}

// __________Pointer receivers___________
// You can declare methods with pointer receivers.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Receiver type has the literal syntax *T for type T. Remember! T cannot itself be a pointer such as *int.
// Methods with pointer receivers can modify the value to which the receiver points
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
  // If we take out the pointer receiver in the Scale method, v is a copy of the Vertex value. Doesn't modify the original
  // methods with pointer receivers take either a value or a pointer as the receiver when they are called
	v.Scale(5) // Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.
  // or this:
	p := &v
	p.Scale(10)

  // methods with value receivers take either a value or a pointer as the receiver when they are called
	var v Vertex
	fmt.Println(v.Abs())
  // the method call p.Abs() is interpreted as (*p).Abs()
	p := &v
	fmt.Println(p.Abs())
}

// implements the same as above but written as a function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10) // functions with a pointer argument must take a pointer
	fmt.Println(Abs(v)) // Functions that take a value argument must take a value of that specific type
}
```

### Interfaces
https://gobyexample.com/interfaces
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.
``` go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser. why? bc main is a funcation!!
	// a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// _____ The empty interface_______
// An empty interface may hold values of any type.
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
// => (<nil>, <nil>)
// => (42, int)
// => (hello, string)

```

### Type assertions
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
If i does not hold a T, the statement will trigger a panic.
``` go
func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	// To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
```

### Stringers
``` go
// One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}
// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
// Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```

## Concurrency

### Goroutines
Lightweight thread manage by Go runtime. Goroutines run in the same address space, so access to shared memory must be synchronized
``` go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
  // `go f()` starts a new goroutine running
  // The evaluation f() happens in the current goroutine and the execution of f happens in the new goroutine. Goroutines run in the same address space, so access to shared memory must be synchronized
	go say("world")
	say("hello")
}
```

### Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-. The data flows in the direction of the arrow. By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
```go
// The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

  // Like maps and slices, channels must be created before use
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

### Buffered Channels
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
``` go
func main() {
  // Provide the buffer length as the second argument to make to initialize a buffered channel:
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

```

### Range and Close
A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after
`v, ok := <-ch`
ok is false if there are no more values to receive and the channel is closed.

``` go
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
  // Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
  // this will receive values from the channel repeatedly until it is closed
	for i := range c {
		fmt.Println(i)
	}
}
```

### Select
The select statement lets a goroutine wait on multiple communication operations.
A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
``` go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

### Default Selection
The default case in a select is run if no other case is ready.


``` go
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
    // Use a default case to try a send or receive without blocking:
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```
