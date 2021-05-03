
# Control structures

## if else 


```go
import "fmt"
```


```go
func basic(number int) string {
    if number > 0 {
        return "positive number"
    } else if number < 0 {
        return "negative number"
    } else {
        return "zero"
    }
}
```


```go
fmt.Println(basic(10))
fmt.Println(basic(-1))
fmt.Println(basic(0))
```

    positive number
    negative number
    zero
    5
    <nil>



```go
func statementCondition(number1 int, number2 int) string {
    // put one statement before the condition
    if total := number1 + number2; total < 100 {
        return "under 100"
    } else if total > 100 {
        return "over 100"
    } else {
        return "100"
    }
}
```


```go
fmt.Println(statementCondition(10, 20))
fmt.Println(statementCondition(90, 20))
fmt.Println(statementCondition(70, 30))
```

    under 100
    over 100
    100
    4
    <nil>


## for loop


```go
func getTotalSum(number int) int {
    total := 0
    for index := 1; index <= number; index ++ {
        total += index
    }
    return total
}


```


```go
fmt.Println(getTotalSum(10))
fmt.Println(getTotalSum(20))
```

    55
    210
    4
    <nil>



```go
//while simulator
func getTotalSumByWhileLoop(number int) int {
    total := 0
    index := 1
    for ; index <= number; {
        total += index
        index ++
    }
    return total
}
```


```go
fmt.Println(getTotalSumByWhileLoop(10))
fmt.Println(getTotalSumByWhileLoop(20))


```

    55
    210
    4
    <nil>


## for with label


```go
//find prime number from 2 to 100
import (
    "fmt";
    "math"
    )

func findPrime(rangeNumber int){
    next:
    for number := 2; number < rangeNumber; number ++ {
        maxNumber := int(math.Sqrt(float64(number))) + 1
        for index := 2; index < maxNumber; index ++ {
            if number % index == 0 {
                continue next
            }
        }
        fmt.Println(number, "is a prime number")
    }
} 
```


```go
findPrime(100)
```

    2 is a prime number
    3 is a prime number
    5 is a prime number
    7 is a prime number
    11 is a prime number
    13 is a prime number
    17 is a prime number
    19 is a prime number
    23 is a prime number
    29 is a prime number
    31 is a prime number
    37 is a prime number
    41 is a prime number
    43 is a prime number
    47 is a prime number
    53 is a prime number
    59 is a prime number
    61 is a prime number
    67 is a prime number
    71 is a prime number
    73 is a prime number
    79 is a prime number
    83 is a prime number
    89 is a prime number
    97 is a prime number



```go

func find5FirstPrime(rangeNumber int){
    count := 0
    next:
    for number := 2; number < rangeNumber; number ++ {
        maxNumber := int(math.Sqrt(float64(number))) + 1
        for index := 2; index < maxNumber; index ++ {
            if number % index == 0 {
                continue next
            }
        }
        fmt.Println(number, "is a prime number")
        count ++
        if count == 5 {
            break next
        }
    }
} 
```


```go
find5FirstPrime(100)
```

    2 is a prime number
    3 is a prime number
    5 is a prime number
    7 is a prime number
    11 is a prime number


# switch


```go
import (
    "fmt";
    "runtime"
)

func basicSwitch(){
    os := runtime.GOOS
    switch os {
    case "linux": 
        fmt.Println("Linux")
    
    case "darwin":
        fmt.Println("MacOS")
    
    default:
        fmt.Println("Other")
    }    
}

basicSwitch()
```

    Linux



```go
func assignStatementSwitch(){
    switch os := runtime.GOOS; os {
    case "linux": 
        fmt.Println("Linux")
    
    case "darwin":
        fmt.Println("MacOS")
    
    default:
        fmt.Println("Other")
    }
}
assignStatementSwitch()

```

    Linux



```go
func compareSwitch(number int){
    switch {
	case number%15 == 0:
		fmt.Println("FIZZBUZZ")
	case number%5 == 0:
		fmt.Println("BUZZ")
	case number%3 == 0:
		fmt.Println("FIZZ")
	default:
		fmt.Println(number)
	}
}

compareSwitch(10)
compareSwitch(9)
compareSwitch(1)
compareSwitch(15)
```

    BUZZ
    FIZZ
    1
    FIZZBUZZ



```go

```
