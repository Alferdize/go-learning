package main

import "fmt"


func base() func() int{
	i := 1
	return func() int {
		j := i*i*i
		i += 1
		return j
	}
}


func main(){
	cube  :=  base()
	fmt.Println(cube())
	fmt.Println(cube())
	fmt.Println(cube())

}
