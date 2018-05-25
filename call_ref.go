package main

import "fmt"


func main(){
	var a int = 9
	b := 10
	fmt.Print("The value of a and b is",a,b)
	swap(&a,&b)
	fmt.Print("\n The values of a and b is ",a,b)
}

func swap(x *int , y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
