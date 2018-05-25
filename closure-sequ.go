package main

import "fmt"

func mul(n int) func(a int) int {
	func(a int) int{
		a = n*a
		return a
	}
}

func main(){
	_mul = mul()
	fmt.Println(_mul(6))
}
