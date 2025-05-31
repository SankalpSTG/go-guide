# Go Tricky Output Based Questions

### Q1. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    funcs := []func(){}
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func() { fmt.Println(i) })
    }
    for _, f := range funcs {
        f()
    }
}
```
### Q2. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    t := s[:2]
    t = append(t, 99)
    fmt.Println("s:", s)
    fmt.Println("t:", t)
}
```
### Q3. What will be the output for the following code
```go
package main

import "fmt"

func f() (x int) {
    defer func() {
        x++
    }()
    return 1
}

func main() {
    fmt.Println(f())
}
```
### Q4. What will be the output for the following code
```go
package main

import "fmt"

type T struct{}

func (t *T) String() string {
    return "I'm T"
}

func main() {
    var t *T = nil
    var i fmt.Stringer = t
    fmt.Println(i == nil)
    fmt.Println(i)
}
```
### Q5. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go func() {
        ch1 <- 1
    }()

    select {
    case x := <-ch1:
        fmt.Println("Received", x)
    case y := <-ch2:
        fmt.Println("Received", y)
    }
}
```
### Q6. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    x := 5
    {
        x := x
        fmt.Println(x)
    }
}
```
### Q7. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    defer fmt.Println("deferred")
    panic("something went wrong")
    fmt.Println("after panic")
}
```
### Q8. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    t := s
    s = append(s[:1], s[2:]...)
    fmt.Println("s:", s)
    fmt.Println("t:", t)
}
```
### Q9. What will be the output for the following code
```go
package main

import "fmt"

type MyStruct struct{}

func (m *MyStruct) Error() string {
    return "error from MyStruct"
}

func main() {
    var m *MyStruct = nil
    var err error = m
    fmt.Println(err == nil)
    fmt.Println(err)
}
```
### Q10. What will be the output for the following code
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)

    select {
    case v := <-ch:
        fmt.Println("Received", v)
    default:
        fmt.Println("No value received")
    }

    time.Sleep(time.Second)
}
```
### Q11. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    for k, v := range m {
        fmt.Println(k, v)
    }
}
```
### Q12. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3}
    var ptrs []*int

    for _, n := range nums {
        ptrs = append(ptrs, &n)
    }

    for _, p := range ptrs {
        fmt.Println(*p)
    }
}
```
### Q13. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    s := "hello"
    b := []byte(s)
    b[0] = 'H'
    fmt.Println(s)
    fmt.Println(string(b))
}
```
### Q14. What will be the output for the following code
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
    time.Sleep(time.Second)
}
```
### Q15. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var i interface{} = 10
    s := i.(string)
    fmt.Println(s)
}
```
### Q16. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    var ptrs []*int

    for _, v := range m {
        ptrs = append(ptrs, &v)
    }

    for _, p := range ptrs {
        fmt.Println(*p)
    }
}
```
### Q17. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4}
    b := a[:2]
    c := a[2:]
    b = append(b, 99)
    fmt.Println("a:", a)
    fmt.Println("b:", b)
    fmt.Println("c:", c)
}
```
### Q18. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        } else {
            fmt.Println("No panic occurred")
        }
    }()
    fmt.Println("Doing fine")
}
```
### Q19. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    const x = 10
    {
        var x = 20
        fmt.Println(x)
    }
    fmt.Println(x)
}
```
### Q20. What will be the output for the following code
```go
package main

import "fmt"

type Person struct {
    name string
}

func main() {
    p1 := Person{name: "Alice"}
    p2 := p1
    p2.name = "Bob"

    fmt.Println(p1.name)
    fmt.Println(p2.name)
}
```
### Q21. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var a []int
    b := []int{}

    fmt.Println(a == nil)
    fmt.Println(b == nil)
}
```
### Q22. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; fmt.Println(i) {
        i++
    }
}
```
### Q23. What will be the output for the following code
```go
package main

import "fmt"

func modifyArr(a [3]int) {
    a[0] = 100
}

func modifySlice(s []int) {
    s[0] = 100
}

func main() {
    arr := [3]int{1, 2, 3}
    slice := []int{1, 2, 3}

    modifyArr(arr)
    modifySlice(slice)

    fmt.Println(arr)
    fmt.Println(slice)
}
```
### Q24. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    res := make([][]int, 0)
    row := []int{}

    for i := 0; i < 3; i++ {
        row = append(row, i)
        res = append(res, row)
    }

    fmt.Println(res)
}
```
### Q25. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    x := 2
    switch x {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println("Default")
    }
}
```
### Q26. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var nums []int

    for _, v := range nums {
        fmt.Println(v)
    }

    fmt.Println("done")
}
```
### Q27. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    ch <- 10
    fmt.Println("done")
}
```
### Q28. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    s := make([]int, 0, 2)
    capBefore := cap(s)

    s = append(s, 1)
    s = append(s, 2)
    capAfter := cap(s)

    s = append(s, 3)
    capFinal := cap(s)

    fmt.Println(capBefore, capAfter, capFinal)
}
```
### Q29. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var i interface{} = nil
    fmt.Println(i == nil)

    var p *int = nil
    i = p
    fmt.Println(i == nil)
}
```
### Q30. What will be the output for the following code
```go
package main

import "fmt"

type A struct{}

func (a A) Foo() {
    fmt.Println("A Foo")
}

type B struct {
    A
}

func main() {
    b := B{}
    b.Foo()
}
```
### Q31. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var s1 []int
    s2 := []int{}

    fmt.Println(len(s1), cap(s1), s1 == nil)
    fmt.Println(len(s2), cap(s2), s2 == nil)
}
```
### Q32. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    x := 5
    defer fmt.Println(x)
    x = 10
}
```
### Q33. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var m1 map[string]int
    m2 := make(map[string]int)

    fmt.Println(m1 == nil)
    fmt.Println(m2 == nil)

    m2["key"] = 1
    // m1["key"] = 1  // Uncommenting this will cause what?
}
```
### Q34. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    s1 := s[:3]
    s2 := s[3:]

    s1 = append(s1, 10)
    s2[0] = 20

    fmt.Println(s)
    fmt.Println(s1)
    fmt.Println(s2)
}
```
### Q35. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var i interface{} = (*int)(nil)

    v, ok := i.(*int)
    fmt.Println(v, ok)
}

```
### Q36. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    m := map[int]string{
        1: "one",
        2: "two",
        3: "three",
    }

    for k := range m {
        fmt.Print(k, " ")
    }
}
```
### Q37. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    for v := range ch {
        fmt.Print(v, " ")
    }
}

```
### Q38. What will be the output for the following code
```go
package main

import (
    "encoding/json"
    "fmt"
)

type Base struct {
    ID int `json:"id"`
}

type User struct {
    Base
    Name string `json:"name"`
}

func main() {
    data := []byte(`{"id":10, "name":"Alice"}`)
    var u User
    json.Unmarshal(data, &u)
    fmt.Println(u.ID, u.Name)
}

```
### Q39. What will be the output for the following code
```go
package main

import "fmt"

const (
    a = iota
    b
    c = 5
    d
    e = iota
)

func main() {
    fmt.Println(a, b, c, d, e)
}
```
### Q40. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    var err error
    fmt.Println(err == nil)

    err = (*struct{})(nil)
    fmt.Println(err == nil)
}
```
### Q41. What will be the output for the following code
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    s1 := s[:3]
    s2 := s1[:4]

    fmt.Println(s2)
}

```