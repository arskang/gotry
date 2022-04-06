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
		Catch: func(e *gotry.Exception) {
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
    err :=  gotry.New{
		Try: func() {
			err := true
			fmt.Println("Start")
			if err {
				panic("Error")
			}
			fmt.Println("End")
		},
        Catch: func(e *gotry.Exception) {
			// Not available
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

3. Code: **.Execute()**
```go
package main

import "github.com/arskang/gotry"

func main() {
    success, err :=  gotry.New{
		Try: func()  gotry.Success {
			err := false
			fmt.Println("Start")
			if err {
				panic("Error")
			}
			return "End"
		},
        Catch: func(e *gotry.Exception) {
			// Not available
		},
        Finally: func() {
			fmt.Println("Finally")
		},
	}.Execute()
	fmt.Println("Error:", err)
	fmt.Println("Success:", success)
}
```

```cmd
Start
Finally
Error: <nil>
Success: End
```