package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("보완한 panic: %v\n", r)
		}
	}()

	fmt.Println("출력된다")
	panic("happening!")
	fmt.Println("출력되지 않는다")
}
