

```go
import "fmt"
```

# Hello by GO


```go
func hello_go(){
    fmt.Println("hello-world by Golang")
}

hello_go()
```

    hello-world by Golang


# Declarations 


```go
var foo int // declaration without initialization
```


```go
var foo int = 42 // declaration with initialization
fmt.Println("foo =", foo)
```

    foo = 42
    9
    <nil>



```go
var foo, bar int = 42, 1302 // declare and init multiple vars at once
fmt.Println("foo =", foo, "type: ", reflect.ValueOf(foo).Kind())
fmt.Println("bar", bar, "type: ", reflect.ValueOf(bar).Kind())
```

    foo = 42 type:  int
    bar 1302 type:  int
    20
    <nil>



```go
var foo = 42 // type omitted, will be inferred 
fmt.Println("foo =", foo, "type: ", reflect.ValueOf(foo).Kind())
```

    foo = 42 type:  int
    20
    <nil>



```go
foo := 42 // shorthand, only in func bodies, omit var keyword, type is always implicit
fmt.Println("foo =", foo, "type: ", reflect.ValueOf(foo).Kind())
```

    foo = 42 type:  int
    20
    <nil>


# Functions


```go
// a simple function
func functionName() {
    fmt.Println("a simple function")
}
```


```go
functionName()
```

    a simple function



```go
// function with parameters (again, types go after identifiers)
func functionName(param1 string, param2 int) {
    fmt.Println("function with parameters (again, types go after identifiers)", param1, param2)
}
```


```go
functionName("vinh", 35)
```

    function with parameters (again, types go after identifiers) vinh 35



```go
// multiple parameters of the same type
func functionName(param1, param2 int) {
    fmt.Println("multiple parameters of the same type", param1, param2)
}

```


```go
functionName(100, 200)
```

    multiple parameters of the same type 100 200



```go
// return type declaration
func functionName() int {
    fmt.Println("return type declaration")
    return 42
}
```


```go
var value = functionName()
fmt.Println(value)
```

    return type declaration
    42
    3
    <nil>



```go
// Can return multiple values at once
func returnMulti() (string, int) {
    fmt.Println("return multiple value")
    return "Vinh", 35
}
```


```go
var name, age = returnMulti()
fmt.Println(name, age)
```

    return multiple value
    Vinh 35
    8
    <nil>



```go
// Return multiple named results simply by return
func returnMulti2() (name string, age int) {
    name = "Vinh"
    age = 35
    // name and age will be returned
    return
}

name, age = returnMulti2()
fmt.Println(name, age)
```

    Vinh 35
    8
    <nil>


## Variadic Functions



```go
func adder(args ...int) int {
    total := 0
    for _, value := range args {
        total += value
    }
    return total
}
```


```go
value1 := adder(1, 2, 3)
value2 := adder(4, 5)
value3 := adder(6, 7, 8, 9)

fmt.Println(value1, value2, value3)
```

    6 9 30
    7
    <nil>



```go

```
