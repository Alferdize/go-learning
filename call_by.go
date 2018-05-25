package main

import "fmt"

func main(){
	var a int = 12
	var b int = 18
	fmt.Print("The value of a and b is",a,b)
	a , b = swap(a,b)
	fmt.Print("The value of a and b is",a,b)
}

func swap(a, b int) (int , int){
	var temp int
	temp = a
	a = b
	b = temp
	return a, b
}
