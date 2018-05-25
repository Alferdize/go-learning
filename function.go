package main

import "fmt"

func main(){
	a := 4
	var b = 7
	fmt.Println("The a and b before swap",a,b)
	a,b = swap(a,b)
	fmt.Print("The a and b before swap%d %d",a,b)

}

func swap(a, b int)(int, int){
	var temp = a
	a = b
	b = temp
	return a, b
}
