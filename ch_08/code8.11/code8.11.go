package code8_11

import "fmt"

type ErrCode int

type MyError struct {
	Code ErrCode
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d", e.Code)
}

//func (e *MyError) Unwrap() error              { return nil }
//func (e *MyError) As(target interface{}) bool { /* 구현 */ }
//func (e *MyError) Is(target error) bool       { /* 구현 */ }
