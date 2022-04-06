package gotry

import (
	"fmt"
	"testing"
)

func Test_block_Go(t *testing.T) {
	New{
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

func Test_block_Error(t *testing.T) {
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
