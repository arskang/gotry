# GOTRY

A simple Try-Catch-Finally with Golang

Credits: https://dzone.com/articles/try-and-catch-in-golang

##### Examples

1. Code: **.Go()**
```go
package main

import "github.com/arskang/gotry"

func main() {
    gotry.New{
		Try: func() {
			err := true
			fmt.Println("Start")
			if err {
				panic("Error")
			}
			fmt.Println("End")
		},
		Catch: func(e *Exception) {
			fmt.Println("Catch:", e.Type, e.Error)
		},
		Finally: func() {
			fmt.Println("Finally")
		},
	}.Go()
}
```

```cmd
Start
Catch: string Error
Finally
```
2. Code: **.Error()**
```go
package main

import "github.com/arskang/gotry"

func main() {
    err := New{
		Try: func() {
			err := true
			fmt.Println("Start")
			if err {
				panic("Error")
			}
			fmt.Println("End")
		},
		Finally: func() {
			fmt.Println("Finally")
		},
	}.Error()
	fmt.Println("Error:", err)
}
```

```cmd
Start
Finally
Error:  &{string Error}
```