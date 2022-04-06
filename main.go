package gotry

import (
	"reflect"
)

// Exception...
type Exception struct {
	Type  reflect.Type
	Error interface{}
}

// New...
// Try: Try to execute the function.
// Catch: Catch the error.
// Finally: Execute to the end.
type New struct {
	Try     func()
	Catch   func(*Exception)
	Finally func()
}

// try...
func (n New) try() *Exception {
	if n.Try != nil {
		n.Try()
	}
	return nil
}

// finally..
func (n New) finally() {
	if n.Finally != nil {
		n.Finally()
	}
}

// Go...
// Exec Try, Catch and Finally.
func (n New) Go() {
	defer n.finally()
	if n.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				n.Catch(&Exception{
					Error: r,
					Type:  reflect.TypeOf(r),
				})
			}
		}()
	}
	n.try()
}

// Error...
// Execute Try and Finally.
// Return an Exception.
func (n New) Error() (e *Exception) {
	defer n.finally()
	defer func() {
		if r := recover(); r != nil {
			e = &Exception{
				Error: r,
				Type:  reflect.TypeOf(r),
			}
		}
	}()
	return n.try()
}
