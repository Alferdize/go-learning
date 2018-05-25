package main

import "fmt"

func main(){
	var a, b = "herf", "freeze"
	fmt.Print(a,b)
	a, b = swap(a,b)
	print(a,b)
}

func swap(a, b string) (string, string){
	var temp = a
	a = b
	b = temp
	return a, b
}
