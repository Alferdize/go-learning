package main

import "fmt"
func max(num1, num2 int)int{
	if num1 > num2{
		return num1
	}else{
		return num2
	}

}
func main(){
	var num1 , num2 = 3, 5
	fmt.Println(max(num1,num2))

}
