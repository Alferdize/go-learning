package main

import "fmt"

func main(){
	var x int;
	var y float32;
	var bit byte;
	x = 10;
	i  := 10.6;
	bit = 56;
	y = 4.5;
	fmt.Print("The x is ",x)
	fmt.Println("\nThe y is ",y)
	fmt.Println("The i is ",i)
	fmt.Println("The bit is ",bit)
	var h , j , k = 2 , 4.5 , "foo"
	fmt.Println(h,j,k)
}
