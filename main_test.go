package gotry

import (
	"fmt"
	"testing"
)

func Test_block_Go(t *testing.T) {
	New{
		Try: func() Success {
			err := true
			fmt.Println("Start")
			if err {
				panic("Error")
			}
			fmt.Println("End")
			return nil
		},
		Catch: func(e *Exception) {
			fmt.Println("Catch:", e.Type, e.Error)
		},
		Finally: func() {
			fmt.Println("Finally")
		},
	}.Go()
}

func Test_block_Error(t *testing.T) {
	err := New{
		Try: func() Success {
			err := true
			fmt.Println("Start")
			if err {
				panic("Error")
			}
			fmt.Println("End")
			return nil
		},
		Catch: func(e *Exception) {
			// Not available
		},
		Finally: func() {
			fmt.Println("Finally")
		},
	}.Error()
	fmt.Println("Error:", err)
}

type Principal struct {
	Sub *Sub
}

type Sub struct {
	Text string
}

func Test_block_Execute(t *testing.T) {
	success, err := New{
		Try: func() Success {
			var p *Principal
			fmt.Println(p.Sub.Text)
			err := false
			fmt.Println("Start")
			if err {
				panic("Error")
			}
			return "End"
		},
		Catch: func(e *Exception) {
			// Not available
		},
		Finally: func() {
			fmt.Println("Finally")
		},
	}.Execute()
	fmt.Println("Error:", err)
	fmt.Println("Success:", success)
}
