package main

import (
	"errors"
	"fmt"
)

var (
	ErrorInF1 error = fmt.Errorf("Error in f1")
	ErrorInF2 error = fmt.Errorf("Error in f2")
	ErrorInF3 error = fmt.Errorf("Error in f3")
	result    string
)

type MyError struct {
	Func string
	Err  error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error in %s: %v", e.Func, e.Err)
}
func (e *MyError) Unwrap() error {
	return e.Err
}
func f1() error {
	err := f2()
	if err != nil {
		return &MyError{Func: "f1", Err: fmt.Errorf("%w:%w", ErrorInF1, err)}
	}
	return nil
}
func f2() error {
	err := f3() //nil
	if err != nil {
		return &MyError{Func: "f2", Err: fmt.Errorf("%w:%w", ErrorInF2, err)}
	}
	return &MyError{Func: "f2", Err: ErrorInF2}
}
func f3() error {
	return nil
}
func findErr(err error) {
	for err != nil {
		var mnr *MyError //mnr - My New Error
		if errors.As(err, &mnr) {
			result = mnr.Func
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
		case interface{ Unwrap() []error }: //wrap error - []error
			for _, e := range x.Unwrap() {
				findErr(e)
			}
			return
		default:
			return
		}
	}
}
func main() {
	err := f1()
	if err != nil {
		findErr(err)
	}
	fmt.Printf("In function: %s\n", result)
}
