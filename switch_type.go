package main 

import "fmt"

func main(){
var x interface {}

switch i := x.(type){
case nil:
fmt.Println("Type of x is ",i)
case int:
fmt.Println("Type of x is int")
case float64:
fmt.Println("Type of x is  float64")
case func(int) float64:
fmt.Println("Type of x is func(int)")
case bool, string:
fmt.Println("Type of x is boolean or string")
default:
fmt.Println("don't know the type")
}

}